import React from "react";
import { Application } from "../../../models/application.model";
import { SubmitHandler, useForm } from "react-hook-form";

const StudentDormApplication: React.FC<{ application?: Application }> = () => {
  const { register, handleSubmit, watch } = useForm<Application>();

  const onSubmit: SubmitHandler<Application> = (data) => {
    console.log(data);
  };
  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col"></form>
    </>
  );
};

export default StudentDormApplication;
