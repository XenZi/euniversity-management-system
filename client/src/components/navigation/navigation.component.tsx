import { useDispatch, useSelector } from "react-redux";
import { Link, useNavigate } from "react-router-dom";
import useLocalStorage from "../../hooks/local-storage.hook";
import { removeUser, setUser } from "../../redux/slices/user.slice";
import { RootState } from "../../redux/store/user.store";
import { useState } from "react";
import { axiosInstance } from "../../services/axios.service";
import { RemoveQuotationMarksFromString } from "../../utils/converter.utils";

const Navigation = () => {
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const [, setUserInLocalStorage, removeLocalStorageValue] = useLocalStorage(
    "user",
    null
  );
  const [, setTokenInLocalStorage, removeTokenFromLocalStorage] =
    useLocalStorage("token", null);
  const user = useSelector((state: RootState) => state.user.user);
  const [isChangeRoleVisible, setChangeRoleVisible] = useState<boolean>(false);

  const navLinks = [
    {
      url: "/home",
      name: "Home",
    },
    {
      url: "/dorm",
      name: "Dorm",
    },
  ];
  const navActions = [
    {
      name: "Logout",
      action: () => {
        removeLocalStorageValue();
        dispatch(removeUser());
        removeTokenFromLocalStorage();
        navigate("/");
      },
    },
    {
      name: "Change Role",
      action: () => {
        setChangeRoleVisible(!isChangeRoleVisible);
      },
    },
  ];

  const changeRole = async (index: number) => {
    await axiosInstance
      .post(
        `/auth/switchRole/${user?.personalIdentificationNumber}/${user?.roles[index]}`
      )
      .then((resp) => {
        console.log(resp.data.data);
        setUserInLocalStorage(resp.data.data.user);
        dispatch(setUser({ ...resp.data.data.user }));
        setTokenInLocalStorage(resp.data.data.token);
      })
      .catch((err) => {
        console.error(err);
      });
  };
  return (
    <div className="bg-papaya-500 w-full p-3">
      <div className="max-w-7xl mx-auto px-8 relative">
        <div className="flex items-center justify-between h-16 ">
          <Link to={"/"}>EUniversity</Link>
          <div className="flex space-x-4">
            {navLinks.map((link, i) => {
              return (
                <Link
                  to={link.url}
                  key={i}
                  className="px-3 py-3 hover:text-teal-500"
                >
                  {link.name}
                </Link>
              );
            })}
            {navActions.map((action, i) => {
              return (
                <button
                  onClick={(e) => {
                    e.preventDefault();
                    action.action();
                  }}
                  key={i}
                >
                  {action.name}
                </button>
              );
            })}
          </div>
        </div>
        <div
          className={`absolute right-0 top-100 w-48 bg-auburn-500 border-auburn-500 ${
            isChangeRoleVisible ? "block" : "hidden"
          }`}
        >
          <ul className="list-style-none">
            {user?.roles.map((role, i) => {
              return (
                <li
                  className="p-3 text-white cursor-pointer hover:bg-auburn-600"
                  key={i}
                  onClick={() => {
                    changeRole(i);
                  }}
                >
                  {role}
                </li>
              );
            })}
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Navigation;
