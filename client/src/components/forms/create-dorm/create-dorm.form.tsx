import { useState } from "react";
import { axiosInstance } from "../../../services/axios.service";
import { castFromApplicationTypeNumberToActualString } from "../../../utils/converter.utils";

interface CreateFormData {
  name: string;
  location: string;
  prices: Prices[];
}

interface Prices {
  applicationType: number;
  price: number;
}
const CreateDormForm = () => {
  const [createFormData, setCreateFormData] = useState<CreateFormData>({
    name: "",
    location: "",
    prices: [
      { applicationType: 1, price: 0 },
      { applicationType: 2, price: 0 },
      { applicationType: 3, price: 0 },
      { applicationType: 4, price: 0 },
    ],
  });

  const onInputChange = (
    e: React.FormEvent<HTMLInputElement>,
    key: keyof CreateFormData,
    priceIndex?: number
  ) => {
    const value = e.currentTarget.value;
    setCreateFormData((prevState) => {
      if (key === "prices" && priceIndex !== undefined) {
        const updatedPrices = [...prevState.prices];
        updatedPrices[priceIndex] = {
          ...updatedPrices[priceIndex],
          price: parseFloat(value) || 0,
        };
        return {
          ...prevState,
          prices: updatedPrices,
        };
      }
      return {
        ...prevState,
        [key]: value,
      };
    });
  };

  const submitForm = async () => {
    await axiosInstance
      .post("/dorm/", createFormData)
      .then((resp) => {
        console.log(resp.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <form action="#" className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">
        Create new dorm
      </h2>
      <input
        type="text"
        name="name"
        id="name"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Dorm name..."
        value={createFormData.name}
        onChange={(e) => onInputChange(e, "name")}
      />
      <input
        type="text"
        name="location"
        id="location"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Dorm location..."
        value={createFormData.location}
        onChange={(e) => onInputChange(e, "location")}
      />
      <p className="font-bold mb-3">Prices</p>
      {createFormData.prices.map((price, index) => (
        <input
          key={index}
          type="number"
          name={`price-${index}`}
          id={`price-${index}`}
          className="mb-3 p-3 border-2 border-battleship-500"
          placeholder={`Price for application ${castFromApplicationTypeNumberToActualString(
            price.applicationType
          )}`}
          onChange={(e) => onInputChange(e, "prices", index)}
        />
      ))}
      <button
        className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
        onClick={(e) => {
          e.preventDefault();
          submitForm();
        }}
      >
        Create Form
      </button>
    </form>
  );
};

export default CreateDormForm;
