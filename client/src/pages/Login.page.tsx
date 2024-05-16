import React, { useEffect, useState } from "react";
import LoginForm from "../components/forms/login.form";
import { useSelector } from "react-redux";
import { RootState } from "../redux/store/user.store";
import { useNavigate } from "react-router-dom";

const LoginPage = () => {
  const user = useSelector((state: RootState) => state.user.user);
  const navigate = useNavigate();
  useEffect(() => {
    if (user) {
      navigate("/home");
    }
  }, [user]);
  return (
    <div className="h-screen bg-papaya-500">
      <div className="max-w-7xl mx-auto w-100 h-full">
        <div className="flex h-full w-100 items-center justify-center">
          <LoginForm />
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
