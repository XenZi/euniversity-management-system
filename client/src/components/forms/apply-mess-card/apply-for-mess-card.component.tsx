import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../../redux/store/store";
import { closeModal } from "../../../redux/slices/modal.slice";
import { axiosInstance } from "../../../services/axios.service";

const ApplyForMess = () => {
  const dispatch = useDispatch();
  const user = useSelector((state: RootState) => state.user.user);

  const [messRooms, setMessRooms] = useState<{ id: string; name: string }[]>(
    []
  );
  const [messRoomId, setMessRoomId] = useState<string>("");
  const [messRoomName, setMessRoomName] = useState<string>("");

  useEffect(() => {
    const fetchMessRooms = async () => {
      try {
        const response = await axiosInstance.get("/food/allMessRooms");
        const fetchedMessRooms = response.data.data;

        // Ensure user and user.personalIdentificationNumber are available
        if (user && user.personalIdentificationNumber) {
          // Filter mess rooms where user's personalIdentificationNumber is not in mess_room_users
          const filteredMessRooms = fetchedMessRooms.filter(
            (messRoom: { mess_room_users: string | string[] }) =>
              !messRoom.mess_room_users.includes(
                user.personalIdentificationNumber
              )
          );

          setMessRooms(filteredMessRooms);

          // If there's only one mess room after filtering, set the messRoomId
          if (filteredMessRooms.length === 1) {
            setMessRoomId(filteredMessRooms[0].id);
            console.log("filterovane menze su", filteredMessRooms[0]);
            setMessRoomName(filteredMessRooms[0].name);
          }

          console.log(filteredMessRooms);
        } else {
          // Handle case where user is null
          console.error("User information is unavailable");
        }
      } catch (error) {
        console.error("Failed to fetch messRooms:", error);
      }
    };

    fetchMessRooms();
  }, [user]);

  const submitForm = async () => {
    if (!user) {
      console.log("User is not available.");
      return;
    }
    const createFoodCardData = {
      student_pin: user.personalIdentificationNumber,
      expires: new Date(Date.now() + 365 * 24 * 60 * 60 * 1000).toISOString(), // Converts date to string in ISO format
      messroom_name: messRoomName,
      mass_room_id: messRoomId,
    };

    console.log(createFoodCardData);

    try {
      console.log(
        axiosInstance.defaults.baseURL + "/food/createFoodCard",
        createFoodCardData
      );
      const resp = await axiosInstance.post(
        "/food/createFoodCard",
        createFoodCardData
      );
      console.log(resp.data.data);
      console.log("createFoodCardData", createFoodCardData);
      var studentPIN = createFoodCardData.student_pin;
      var messID = createFoodCardData.mass_room_id;
      console.log("Mess room id je", messRoomId);

      const response = await axiosInstance.put(
        `/food/updateMessUsers/${messID}/${studentPIN}`
      );

      console.log(response);
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
        className="bg-light p-4 rounded shadow-sm"
      >
        <div className="mb-3">
          <label htmlFor="messRoom" className="form-label">
            Select Mess Room:
          </label>
          <select
            id="messRoom"
            className="mb-3 p-3 border-2 border-battleship-500"
            value={messRoomId}
            onChange={(e) => {
              console.log("Selected messRoomId:", e.target.value); // Debugging line
              setMessRoomId(e.target.value); // Set the selected value
            }}
          >
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
          className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
        >
          Submit
        </button>
      </form>
    </div>
  );
};

export default ApplyForMess;
