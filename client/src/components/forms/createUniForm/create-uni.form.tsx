import axios from "axios";
import { useState } from "react";
import { axiosInstance } from "../../../services/axios.service";

interface CreateFormData {
  name: string;
  address: string;
}

const CreateUniForm: React.FC = () => {
  const [createFormData, setCreateFormData] = useState<CreateFormData>({
    name: "",
    address: ""
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setCreateFormData(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await axiosInstance.post("/university/", createFormData);
      console.log(response.data.data);
    } catch (error) {
      console.log(error);
    }
    console.log('Form submitted:', createFormData);
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">
        Create New University
      </h2>
      <div className="mb-4">
        <label htmlFor="name" className="block text-sm font-medium text-gray-700">
          Name
        </label>
        <input
          type="text"
          name="name"
          id="name"
          value={createFormData.name}
          onChange={handleChange}
          className="mb-3 p-3 border-2 border-battleship-500"
          required
        />
      </div>
      <div className="mb-4">
        <label htmlFor="address" className="block text-sm font-medium text-gray-700">
          Address
        </label>
        <input
          type="text"
          name="address"
          id="address"
          value={createFormData.address}
          onChange={handleChange}
          className="mb-3 p-3 border-2 border-battleship-500"
          required
        />
      </div>
      <button
        type="submit"
        className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
      >
        Create University
      </button>
    </form>
  );
};

export default CreateUniForm;
