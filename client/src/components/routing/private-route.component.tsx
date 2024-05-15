import { useSelector } from "react-redux";
import { Navigate } from "react-router-dom";
import { RootState } from "../../redux/store/user.store";

interface PrivateRouteProps {
  Component: React.ComponentType<unknown>;
  requiredRole: string;
}

const PrivateRoute: React.FC<PrivateRouteProps> = ({
  Component,
  requiredRole,
  ...rest
}) => {
  const user = useSelector((state: RootState) => state.user.user);
  return user && user.roles[0] == requiredRole ? (
    <Component {...rest} />
  ) : (
    <Navigate to="/" replace />
  );
};

export default PrivateRoute;
