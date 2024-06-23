import { useEffect, useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { axiosInstance } from "../../../services/axios.service";
export interface CreateRoomData {
  dormID: string;
  squareFoot: number;
  numberOfBeds: number;
  toalet: number;
}

export enum ToaletType {
  RoomShared = 0,
  FloorShared = 1,
  RoomBased = 2,
}

interface DormData {
  id: string;
  name: string;
  location: string;
}

const CreateRoomForm = () => {
  const { register, handleSubmit } = useForm<CreateRoomData>();
  const onSubmit: SubmitHandler<CreateRoomData> = (data) => {
    console.log(data);
  };
  const [dormData, setDormData] = useState<DormData[]>();

  useEffect(() => {
    loadData();
  }, []);
  const loadData = async () => {
    await axiosInstance.get(`/dorm/all`).then((data) => {
      setDormData(data.data.data);
    });
  };
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">
        Create new room for dorm
      </h2>
      <div className="w-full">
        <label>Choose dorm</label>
        <select
          id="color-select"
          className="mb-3 p-3 border-2 border-battleship-500 w-full"
          {...register("dormID", {
            required: true,
          })}
        >
          {dormData?.map((dorm) => (
            <option value={dorm.id} key={dorm.id}>
              {dorm.name} - {dorm.name}
            </option>
          ))}
        </select>
      </div>
      <input
        type="number"
        id="squareFoot"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Square foot of room..."
        {...register("squareFoot", { required: true, min: 1 })}
      />
      <input
        type="number"
        id="numberOfBeds"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Number of beds in room..."
        {...register("numberOfBeds", { required: true, min: 1 })}
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
          Create Form
        </button>
      </div>
    </form>
  );
};

export default CreateRoomForm;
