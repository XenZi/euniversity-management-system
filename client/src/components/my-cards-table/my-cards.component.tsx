import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useModalContext } from "../../context/modal.context";
import { axiosInstance } from "../../services/axios.service";
import { closeModal, setModalOpen } from "../../redux/slices/modal.slice";
import { FoodCard } from "../../models/foodCard.model";
import { RootState } from "../../redux/store/store";

const MyFoodCardTable: React.FC = () => {
  const [foodCards, setFoodCards] = useState<FoodCard[]>([]);
  const dispatch = useDispatch();
  const { setContent } = useModalContext();

  // Get user data from Redux store
  const user = useSelector((state: RootState) => state.user.user);

  useEffect(() => {
    if (!user || !user.personalIdentificationNumber) {
      // Exit early if user or user.personalIdentificationNumber is null
      return;
    }
    axiosInstance
      .get("/food/allFoodCards")
      .then((data) => {
        // Filter food cards where student_pin matches the user's personalIdentificationNumber
        const filteredFoodCards = data.data.data.filter(
          (card: FoodCard) =>
            card.student_pin === user.personalIdentificationNumber
        );
        setFoodCards(filteredFoodCards);
      })
      .catch((err) => console.log(err));
  }, [user?.personalIdentificationNumber]);

  return (
    <div className="overflow-x-auto">
      <table className="min-w-full bg-white">
        <thead>
          <tr>
            <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
              Card Name
            </th>
            <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
              Expires
            </th>
            <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
              Balance
            </th>
            <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
              Used point
            </th>
          </tr>
        </thead>
        <tbody>
          {foodCards && foodCards.length > 0 ? (
            foodCards.map((card, index) => (
              <tr key={index}>
                <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                  {card.messroom_name}
                </td>
                <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                  {card.expires}
                </td>
                <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                  {card.balance}
                </td>
                <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                  {card.used_point}
                </td>
              </tr>
            ))
          ) : (
            <p className="my-2">No Food Cards available...</p>
          )}
        </tbody>
      </table>
    </div>
  );
};

export default MyFoodCardTable;
