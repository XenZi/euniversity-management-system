import React, { useState } from "react";
import LoginForm from "../components/forms/login.form";

const LoginPage = () => {
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
