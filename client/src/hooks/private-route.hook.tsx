// import { useLocation, Navigate, Location } from 'react-router-dom';
// import { FC, ReactNode } from 'react';

// interface RequireAuthProps {
//   children: ReactNode;
// }

// const RequireAuth: FC<RequireAuthProps> = ({ children }) => {
//   const auth = useAuth();
//   const location = useLocation();

//   if (!auth.user) {
//     return <Navigate to="/login" state={{ from: location }} replace />;
//   }

//   return <>{children}</>;
// };

// export default RequireAuth;
