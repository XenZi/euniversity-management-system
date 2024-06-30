import React from "react";
import { Dorm } from "../../../models/dorm.model";
import {
  Controller,
  SubmitHandler,
  useFieldArray,
  useForm,
} from "react-hook-form";
import { axiosInstance } from "../../../services/axios.service";

const EditDormForm: React.FC<{ dorm: Dorm }> = ({ dorm }) => {
  const { control, handleSubmit, register } = useForm<Dorm>({
    defaultValues: {
      name: dorm.name,
      location: dorm.location,
      prices: dorm.prices,
    },
  });

  const { fields } = useFieldArray({
    control,
    name: "prices",
  });

  const onSubmit: SubmitHandler<Dorm> = (data) => {
    data.id = dorm.id;
    axiosInstance
      .put(`dorm/${dorm.id}`, { ...data })
      .then((data) => console.log(data))
      .catch((err) => console.log(err));
    console.log(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">
        Edit dorm data
      </h2>
      <input
        type="text"
        id="name"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Dorm name..."
        {...register("name")}
      />
      <input
        type="text"
        id="location"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Dorm location..."
        {...register("location")}
      />
      <p className="font-bold mb-3">Prices</p>
      {fields.map((field, index) => (
        <Controller
          name={`prices.${index}.price`}
          control={control}
          key={field.id}
          render={({ field }) => (
            <input
              type="number"
              id={`price-${index}`}
              className="mb-3 p-3 border-2 border-battleship-500"
              placeholder={`Price for application type ${dorm.prices[index].applicationType}`}
              {...field}
              value={field.value as number}
              onChange={(e) => field.onChange(e.target.valueAsNumber)}
            />
          )}
        />
      ))}
      <button
        className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
        type="submit"
      >
        Edit dorm
      </button>
    </form>
  );
};

export default EditDormForm;
