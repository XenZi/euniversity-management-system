import React, { useEffect, useState } from "react";
import { axiosInstance } from "../../services/axios.service";
import { useDispatch } from "react-redux";
import { setUser } from "../../redux/slices/user.slice";
import { useNavigate } from "react-router-dom";
interface LoginFormData {
  email: string;
  password: string;
}
const LoginForm = () => {
  const [loginForm, setLoginForm] = useState<LoginFormData>({
    email: "",
    password: "",
  });
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const onInputChange = (e: React.FormEvent<HTMLInputElement>, key: string) => {
    const copyOfFormData = { ...loginForm };
    copyOfFormData[key as keyof LoginFormData] = e.currentTarget.value;
    setLoginForm(copyOfFormData);
  };

  const sendRequestForLogin = async () => {
    const sendData = await axiosInstance
      .post("/auth/login", loginForm)
      .then((resp) => {
        dispatch(setUser({ ...resp.data.data.user }));
        navigate("/home");
      })
      .catch((err) => {
        console.log(err.response.data);
      });
  };

  useEffect(() => {}, [loginForm]);

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
      <button
        className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
        onClick={(e) => {
          e.preventDefault();
          sendRequestForLogin();
        }}
      >
        Login
      </button>
    </form>
  );
};

export default LoginForm;
