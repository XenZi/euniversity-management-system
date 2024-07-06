import { useSelector } from "react-redux";
import { RootState } from "../../../redux/store/store";

const ProfilePage = () => {
    const user = useSelector((state: RootState) => state.user.user);

    // Check if user exists and contains necessary properties
    if (!user) {
        return null; // Or render a loading indicator or handle the case where user is not available
    }

    return (
        <>
            <div>Profile</div>
            <p>Full Name: {user.fullName}</p>
            <p>PIN: {user.personalIdentificationNumber}</p>
            <p>Citizenship: {user.citizenship}</p>
            <p>Gender: {user.gender}</p>
            <p>IDCard Number: {user.identityCardNumber}</p>
        </>
    );
};

export default ProfilePage;
