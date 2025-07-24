import React from "react";
import { Controller, SubmitHandler, useForm } from "react-hook-form";
import { axiosInstance } from "../../../services/axios.service";
import { useDispatch } from "react-redux";
import { closeModal } from "../../../redux/slices/modal.slice";
import { Mess } from "../../../models/mess.model";

interface RatingFormData {
  rating: number;
}

const RatingForm: React.FC<{ mess: Mess }> = ({ mess }) => {
  const { control, handleSubmit, register } = useForm<Mess>({
    defaultValues: {
      rating: 0, // Initial value for rating
    },
  });
  const dispatch = useDispatch();

  const onSubmit: SubmitHandler<Mess> = (data) => {
    axiosInstance
      .put(`food/updateMessRating/${mess.id}/${data.rating}`)
      .then((response) => console.log(response))
      .catch((err) => console.log(err));
    console.log(data);
    console.log("Proslijedjene vrijednosti su", mess.id, data.rating);
    dispatch(closeModal());
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">Rate the Mess</h2>
      <div className="mb-3">
        <h3 className="text-xl font-semibold mb-2">Rating (1 to 5)</h3>
        <div className="flex space-x-3 justify-center">
          {[1, 2, 3, 4, 5].map((value) => (
            <label key={value} className="flex items-center space-x-1">
              <input
                type="radio"
                value={value}
                {...register("rating")}
                className="form-radio h-5 w-5 text-auburn-500 focus:ring-auburn-500"
              />
              <span className="text-sm font-medium text-gray-700">{value}</span>
            </label>
          ))}
        </div>
      </div>
      <button
        className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
        type="submit"
      >
        Submit Rating
      </button>
    </form>
  );
};

export default RatingForm;
