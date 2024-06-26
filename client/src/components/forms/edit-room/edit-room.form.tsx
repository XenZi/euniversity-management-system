/* eslint-disable @typescript-eslint/no-unused-vars */
import { useForm, SubmitHandler } from "react-hook-form";
import { axiosInstance } from "../../../services/axios.service";
import { Room, ToaletType } from "../../../models/room.model";

const EditRoomForm: React.FC<{ room: Room }> = ({ room }) => {
  const { control, handleSubmit, register } = useForm<Room>({
    defaultValues: {
      squareFoot: room.squareFoot,
      numberOfBeds: room.numberOfBeds,
      toalet: room.toalet,
    },
  });
  const onSubmit: SubmitHandler<Room> = (data) => {
    data.id = room.id;
    axiosInstance
      .put(`/dorm/room/${room.id}`, { ...data })
      .then((data) => console.log(data))
      .catch((err) => console.log(err.response.data));
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">
        Edit dorm room
      </h2>
      <input
        type="number"
        id="squareFoot"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Square foot of room..."
        {...register("squareFoot", {
          required: true,
          min: 1,
          valueAsNumber: true,
        })}
      />
      <input
        type="number"
        id="numberOfBeds"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Number of beds in room..."
        {...register("numberOfBeds", {
          required: true,
          min: 1,
          valueAsNumber: true,
        })}
      />
      <div className="w-full">
        <label
          htmlFor="color-select"
          className="block text-sm font-medium text-gray-700"
        >
          Select Toalet Type
        </label>
        <select
          id="color-select"
          className="mb-3 p-3 border-2 border-battleship-500 w-full"
          {...register("toalet", {
            required: true,
            valueAsNumber: true,
          })}
        >
          {Object.entries(ToaletType)
            .filter(([key]) => isNaN(Number(key)))
            .map(([key, value]) => (
              <option key={key} value={value}>
                {key}
              </option>
            ))}
        </select>
        <button
          className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white w-full"
          type="submit"
        >
          Edit Room
        </button>
      </div>
    </form>
  );
};

export default EditRoomForm;
