import React from "react";
import { Mess } from "../../../models/mess.model";
import { Controller, SubmitHandler, useForm } from "react-hook-form";
import { axiosInstance } from "../../../services/axios.service";
import { useDispatch } from "react-redux";
import { closeModal } from "../../../redux/slices/modal.slice";

const EditMessForm: React.FC<{ mess: Mess }> = ({ mess }) => {
  const { control, handleSubmit, register } = useForm<Mess>({
    defaultValues: {
      name: mess.name,
      location: mess.location,
      capacity: mess.capacity,
    },
  });
  const dispatch = useDispatch();

  const onSubmit: SubmitHandler<Mess> = (data) => {
    data.id = mess.id;
    axiosInstance
      .put(`food/updateMess/${mess.id}`, { ...data })
      .then((response) => console.log(response))
      .catch((err) => console.log(err));
    console.log(data);
    dispatch(closeModal());
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">
        Edit Mess Data
      </h2>
      <input
        type="text"
        id="name"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Mess name..."
        {...register("name")}
      />
      <input
        type="text"
        id="location"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Mess location..."
        {...register("location")}
      />
      <Controller
        name="capacity"
        control={control}
        render={({ field }) => (
          <input
            type="number"
            id="capacity"
            className="mb-3 p-3 border-2 border-battleship-500"
            placeholder="Mess capacity..."
            {...field}
            value={field.value as number}
            onChange={(e) => field.onChange(e.target.valueAsNumber)}
          />
        )}
      />
      <button
        className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
        type="submit"
      >
        Edit Mess
      </button>
    </form>
  );
};

export default EditMessForm;
