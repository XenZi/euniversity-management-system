import { Navigate } from "react-router-dom";
import useLocalStorage from "../../hooks/local-storage.hook";
import { User } from "../../models/user.model";

interface PrivateRouteProps {
  Component: React.ComponentType<unknown>;
}

const PrivateRoute: React.FC<PrivateRouteProps> = ({ Component, ...rest }) => {
  const [userFromLocalStorage, setUserInLocalStorage] = useLocalStorage(
    "user",
    null
  );
  const user = userFromLocalStorage
    ? (userFromLocalStorage as unknown as User)
    : null;
  return user ? <Component {...rest} /> : <Navigate to="/" replace />;
};

export default PrivateRoute;
