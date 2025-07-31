import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../../redux/store/store";
import { closeModal } from "../../../redux/slices/modal.slice";
import { axiosInstance } from "../../../services/axios.service";

const ApplyForMess: React.FC = () => {
  const dispatch = useDispatch();
  const user = useSelector((state: RootState) => state.user.user);

  const [messRooms, setMessRooms] = useState<{ id: string; name: string }[]>(
    []
  );
  const [messRoomId, setMessRoomId] = useState<string>("");

  useEffect(() => {
    const fetchMessRooms = async () => {
      try {
        const response = await axiosInstance.get("/food/allMessRooms");
        const fetchedMessRooms = response.data.data;

        if (user && user.personalIdentificationNumber) {
          // Filter out mess rooms where the user already exists in mess_room_users
          const filteredMessRooms = fetchedMessRooms.filter(
            (messRoom: { mess_room_users: string[] | string }) => {
              const users = Array.isArray(messRoom.mess_room_users)
                ? messRoom.mess_room_users
                : [messRoom.mess_room_users];
              return !users.includes(user.personalIdentificationNumber);
            }
          );

          setMessRooms(filteredMessRooms);

          if (filteredMessRooms.length === 1) {
            setMessRoomId(filteredMessRooms[0].id);
          }
        } else {
          console.error("User information is unavailable");
        }
      } catch (error) {
        console.error("Failed to fetch messRooms:", error);
      }
    };

    fetchMessRooms();
  }, [user]);

  const submitForm = async () => {
    if (!user?.personalIdentificationNumber) {
      console.log("User is not available.");
      return;
    }
    if (!messRoomId) {
      console.log("Mess room not selected.");
      return;
    }

    const selected = messRooms.find((m) => m.id === messRoomId);
    const createFoodCardData = {
      student_pin: user.personalIdentificationNumber,
      expires: new Date(Date.now() + 365 * 24 * 60 * 60 * 1000).toISOString(),
      messroom_name: selected?.name || "",
      mass_room_id: messRoomId,
    };

    try {
      const resp = await axiosInstance.post(
        "/food/createFoodCard",
        createFoodCardData
      );
      console.log("Created food card:", resp.data.data);

      await axiosInstance.put(
        `/food/updateMessUsers/${encodeURIComponent(
          messRoomId
        )}/${encodeURIComponent(createFoodCardData.student_pin)}`
      );

      dispatch(closeModal());
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div className="container mt-4">
      <form
        onSubmit={(e) => {
          e.preventDefault();
          submitForm();
        }}
        className="bg-light p-4 rounded shadow-sm max-w-lg mx-auto"
      >
        <div className="mb-3">
          <label htmlFor="messRoom" className="form-label font-medium">
            Select Mess Room:
          </label>
          <select
            id="messRoom"
            className="w-full mb-3 p-3 border-2 rounded"
            value={messRoomId}
            onChange={(e) => {
              setMessRoomId(e.target.value);
            }}
          >
            <option value="">-- pick a mess --</option>
            {messRooms.length > 0 ? (
              messRooms.map((mess) => (
                <option key={mess.id} value={mess.id}>
                  {mess.name}
                </option>
              ))
            ) : (
              <option value="">No mess rooms available</option>
            )}
          </select>
        </div>

        <button
          type="submit"
          disabled={!messRoomId}
          className="w-full border bg-auburn-500 font-semibold py-2 px-4 rounded text-white hover:bg-auburn-600 disabled:opacity-50"
        >
          Submit
        </button>
      </form>
    </div>
  );
};

export default ApplyForMess;
