import React, { useEffect, useState } from "react";

interface LoginFormData {
  email: string;
  password: string;
}
const LoginForm = () => {
  const [loginForm, setLoginForm] = useState<LoginFormData>({
    email: "",
    password: "",
  });

  const onInputChange = (e: React.FormEvent<HTMLInputElement>, key: string) => {
    const copyOfFormData = { ...loginForm };
    copyOfFormData[key as keyof LoginFormData] = e.currentTarget.value;
    setLoginForm(copyOfFormData);
  };

  useEffect(() => {
    console.log(loginForm);
  }, [loginForm]);

  return (
    <form action="#" className="flex flex-col">
      <input
        type="text"
        name="email"
        id="email"
        className="mb-3 p-3"
        placeholder="Email..."
        onInput={(e) => {
          onInputChange(e, "email");
        }}
      />
      <input
        type="password"
        name="password"
        id="password"
        className="mb-3 p-3"
        placeholder="Password..."
        onInput={(e) => {
          onInputChange(e, "password");
        }}
      />
      <button className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white">
        Login
      </button>
    </form>
  );
};

export default LoginForm;
