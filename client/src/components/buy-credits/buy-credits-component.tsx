import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../redux/store/store";
import { useEffect, useState } from "react";
import { axiosInstance } from "../../services/axios.service";
import { closeModal } from "../../redux/slices/modal.slice";

const BuyCredits = () => {
  const dispatch = useDispatch();
  const user = useSelector((state: RootState) => state.user.user);

  const [messCards, setMessCards] = useState<
    { id: string; messroom_name: string }[]
  >([]);
  const [messCardId, setMessCardId] = useState<string>("");
  const [messCardName, setMessCardName] = useState<string>("");
  const [formData, setFormData] = useState({
    credit_card_number: "",
    name: "",
    cvv: "",
    amount: "",
  });

  useEffect(() => {
    const fetchMessCards = async () => {
      try {
        if (user && user.personalIdentificationNumber) {
          const response = await axiosInstance.get("/food/allFoodCards", {
            params: { student_pin: user.personalIdentificationNumber },
          });
          const fetchedMessCards = response.data.data;
          setMessCards(fetchedMessCards);

          // Auto-select if only one mess card is available
          if (fetchedMessCards.length === 1) {
            setMessCardId(fetchedMessCards[0].id);
            setMessCardName(fetchedMessCards[0].messroom_name);
          }

          console.log(fetchedMessCards);
        } else {
          console.error("User information is unavailable");
        }
      } catch (error) {
        console.error("Failed to fetch mess cards:", error);
      }
    };

    fetchMessCards();
  }, [user]);

  const submitForm = async () => {
    if (!user) {
      console.log("User is not available.");
      return;
    }

    const buyCreditsData = {
      ...formData,
      food_card_id: messCardId,
      messcard_name: messCardName,
    };

    console.log(buyCreditsData);

    try {
      console.log(
        axiosInstance.defaults.baseURL + "/food/createPayment",
        buyCreditsData
      );
      const resp = await axiosInstance.post(
        "/food/createPayment",
        buyCreditsData
      );
      console.log(resp.data);
      dispatch(closeModal());
    } catch (err) {
      console.log(err);
    }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleMessCardChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedCard = messCards.find((card) => card.id === e.target.value);
    if (selectedCard) {
      setMessCardId(selectedCard.id);
      setMessCardName(selectedCard.messroom_name);
    }
  };

  return (
    <form action="#" className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">Buy Credits</h2>

      {/* Credit Card Number */}
      <input
        type="text"
        name="credit_card_number"
        id="credit_card_number"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Credit Card Number..."
        value={formData.credit_card_number}
        onChange={handleInputChange}
        required
      />

      {/* Name on Card */}
      <input
        type="text"
        name="name"
        id="name"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Name on Card..."
        value={formData.name}
        onChange={handleInputChange}
        required
      />

      {/* CVV */}
      <input
        type="text"
        name="cvv"
        id="cvv"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="CVV..."
        value={formData.cvv}
        onChange={handleInputChange}
        required
      />

      {/* Amount */}
      <input
        type="number"
        name="amount"
        id="amount"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Amount..."
        value={formData.amount}
        onChange={handleInputChange}
        required
      />

      {/* Mess Card Selection */}
      <select
        id="messCard"
        className="mb-3 p-3 border-2 border-battleship-500"
        value={messCardId}
        onChange={handleMessCardChange}
        required
      >
        {messCards.length > 0 ? (
          messCards.map((card) => (
            <option key={card.id} value={card.id}>
              {card.messroom_name}
            </option>
          ))
        ) : (
          <option value="">No mess cards available</option>
        )}
      </select>

      {/* Display selected Mess Card Name */}
      {messCardName && (
        <div className="mb-3">
          <strong>Selected Mess Card Name: </strong>
          <span>{messCardName}</span>
        </div>
      )}

      {/* Submit Button */}
      <button
        className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
        onClick={(e) => {
          e.preventDefault();
          submitForm();
        }}
      >
        Buy Credits
      </button>
    </form>
  );
};

export default BuyCredits;
