import React, { useEffect, useState } from "react";
import { Admission } from "../../../models/admission.model";
import { SubmitHandler, useForm } from "react-hook-form";
import { Dorm } from "../../../models/dorm.model";
import { axiosInstance } from "../../../services/axios.service";

const Admissions: React.FC<{ admission?: Admission }> = ({ admission }) => {
  const [loadedDorms, setLoadedDorms] = useState<Dorm[]>([]);

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Admission>({
    defaultValues: {
      id: admission?.id ?? "",
      dormID: admission?.dormID ?? "",
      start: admission?.start ?? "",
      end: admission?.end ?? "",
    },
  });
  const startDate = watch("start");

  const onSubmit: SubmitHandler<Admission> = (data) => {
    if (admission == null || admission == undefined) {
      axiosInstance
        .post("/dorm/admissions", {
          ...data,
        })
        .then((data) => {
          console.log(data.data.data);
        })
        .catch((err) => {
          console.log(err);
        });
      return;
    }

    axiosInstance
      .put(`/dorm/admissions/${data.id}`, {
        ...data,
      })
      .then((data) => {
        console.log(data.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    axiosInstance.get(`/dorm/all`).then((data) => {
      setLoadedDorms(data.data.data);
    });
  }, []);

  useEffect(() => {
    console.log(loadedDorms);
  }, [loadedDorms]);

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col">
        <h2 className="text-center text-3xl font-semibold mb-3">
          {admission ? "Edit admission" : "Create admission"}
        </h2>
        <div className="w-full">
          <label>Choose dorm</label>
          <select
            id="color-select"
            className="mb-3 p-3 border-2 border-battleship-500 w-full"
            {...register("dormID", {
              required: true,
            })}
            disabled={admission ? true : false}
          >
            <option value="" disabled selected>
              Select an option
            </option>
            {loadedDorms?.map((dorm) => (
              <option value={dorm.id} key={dorm.id}>
                {dorm.name} - {dorm.name}
              </option>
            ))}
          </select>
        </div>
        <input
          type="date"
          id="name"
          className="mb-3 p-3 border-2 border-battleship-500"
          placeholder="Start date..."
          {...register("start", {
            required: "Start date is required",
            validate: (value) => {
              const currentDate = new Date().toISOString().split("T")[0];
              return value >= currentDate || "Start date cannot be in the past";
            },
          })}
        />
        <input
          type="date"
          id="location"
          className="mb-3 p-3 border-2 border-battleship-500"
          placeholder="Dorm location..."
          {...register("end", {
            required: "End date is required",
            validate: (value) => {
              return (
                value > startDate || "End date must be after the start date"
              );
            },
          })}
        />
        {errors.start && <p>{errors.start.message}</p>}
        <button
          className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
          type="submit"
        >
          {admission ? "Edit admission" : "Create admission"}
        </button>
      </form>
    </>
  );
};

export default Admissions;
