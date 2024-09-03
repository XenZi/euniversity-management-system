import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import axios from "axios";
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

  useEffect(() => {
    const fetchMessRooms = async () => {
      try {
        const response = await axiosInstance.get("/food/allMessRooms");
        setMessRooms(response.data.data);
      } catch (error) {
        console.error("Failed to fetch messRooms:", error);
      }
    };

    fetchMessRooms();
  }, []);

  const submitForm = async () => {
    if (!user) {
      console.log("User is not available.");
      return;
    }
    const createFoodCardData = {
      student_pin: user.personalIdentificationNumber,
      expires: new Date(Date.now() + 365 * 24 * 60 * 60 * 1000).toISOString(), // Converts date to string in ISO format
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
      var studentPIN = createFoodCardData.student_pin;
      var messID = createFoodCardData.mass_room_id;

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
            onChange={(e) => setMessRoomId(e.target.value)}
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
