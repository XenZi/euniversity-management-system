import { useState } from "react";
import { axiosInstance } from "../../../services/axios.service";
import { useDispatch } from "react-redux";
import { closeModal } from "../../../redux/slices/modal.slice";

interface CreateMessFormData {
  name: string;
  location: string;
  capacity: number;
}

const CreateMessForm = () => {
  const [createMessFormData, setCreateMessFormData] =
    useState<CreateMessFormData>({
      name: "",
      location: "",
      capacity: 0,
    });

  const dispatch = useDispatch();

  const onInputChange = (
    e: React.FormEvent<HTMLInputElement>,
    key: keyof CreateMessFormData
  ) => {
    const value =
      key === "capacity"
        ? parseInt(e.currentTarget.value) || 0
        : e.currentTarget.value;
    setCreateMessFormData((prevState) => ({
      ...prevState,
      [key]: value,
    }));
  };

  const submitForm = async () => {
    await axiosInstance
      .post("/food/createMessRoom", createMessFormData)
      .then((resp) => {
        console.log(resp.data.data);
        dispatch(closeModal());
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <form action="#" className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">
        Create new Mess
      </h2>
      <input
        type="text"
        name="name"
        id="name"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Mess name..."
        value={createMessFormData.name}
        onChange={(e) => onInputChange(e, "name")}
      />
      <input
        type="text"
        name="location"
        id="location"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Mess location..."
        value={createMessFormData.location}
        onChange={(e) => onInputChange(e, "location")}
      />
      <input
        type="number"
        name="capacity"
        id="capacity"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Capacity..."
        value={createMessFormData.capacity}
        onChange={(e) => onInputChange(e, "capacity")}
      />
      <button
        className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
        onClick={(e) => {
          e.preventDefault();
          submitForm();
        }}
      >
        Create Mess
      </button>
    </form>
  );
};

export default CreateMessForm;
