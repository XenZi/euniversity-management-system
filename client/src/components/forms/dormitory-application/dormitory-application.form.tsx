import { useEffect, useState } from "react";
import { Dorm } from "../../../models/dorm.model";
import { Admission } from "../../../models/admission.model";
import { Application } from "../../../models/application.model";
import { SubmitHandler, useForm } from "react-hook-form";
import { axiosInstance } from "../../../services/axios.service";
import { ApplicationType } from "../../../models/enum";
import { RootState } from "../../../redux/store/store";
import { useDispatch, useSelector } from "react-redux";
import { User } from "../../../models/user.model";
import { closeModal } from "../../../redux/slices/modal.slice";

const DormitoryApplication = () => {
  const [dorms, setDorms] = useState<Dorm[]>([]);
  const [admissions, setAdmissions] = useState<Admission[]>([]);
  const { register, handleSubmit } = useForm<Application>();
  const dispatch = useDispatch();
  const user = useSelector((state: RootState) => state.user.user);
  useEffect(() => {
    axiosInstance
      .get(`/dorm/all`)
      .then((data) => {
        setDorms(data.data.data);
      })
      .catch((err) => console.log(err));
    axiosInstance
      .get(`/dorm/admissions/all`)
      .then((data) => {
        setAdmissions(data.data.data);
      })
      .catch((err) => console.log(err));
  }, []);

  const findDormBasedOnDormID = (dormID: string): Dorm | undefined => {
    const foundDorm = dorms.find((dorm) => dorm.id == dormID);
    return foundDorm;
  };

  const onSubmit: SubmitHandler<Application> = (data) => {
    data.student = user as User;
    axiosInstance.post("/dorm/applications", { ...data }).then((data) => {
      console.log(data.data.data);
      dispatch(closeModal());
    });
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">
        Send application for dorm
      </h2>
      <div className="w-full">
        <label>Choose admission</label>
        <select
          id="color-select"
          className="mb-3 p-3 border-2 border-battleship-500 w-full"
          {...register("dormitoryAdmissionsID", {
            required: true,
          })}
        >
          {admissions?.map((admission) => {
            const foundDorm = findDormBasedOnDormID(admission.dormID);
            return (
              <>
                <option value={admission.id} key={admission.id}>
                  {foundDorm?.name}, {foundDorm?.location} - {admission.start} -{" "}
                  {admission.end}
                </option>
              </>
            );
          })}
        </select>
      </div>
      <div className="w-full">
        <label>Choose application type</label>
        <select
          id="color-select"
          className="mb-3 p-3 border-2 border-battleship-500 w-full"
          {...register("applicationType", {
            required: true,
            valueAsNumber: true,
          })}
        >
          {Object.entries(ApplicationType)
            .filter(([key]) => isNaN(Number(key)))
            .map(([key, value]) => (
              <option key={key} value={value}>
                {key}
              </option>
            ))}
        </select>
      </div>
      <button
        className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white w-full"
        type="submit"
      >
        Send application
      </button>
    </form>
  );
};

export default DormitoryApplication;
