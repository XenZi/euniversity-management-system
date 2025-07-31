import { useEffect, useMemo, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { axiosInstance } from "../../services/axios.service";
import { RootState } from "../../redux/store/store";
import { closeModal } from "../../redux/slices/modal.slice";

type FoodCard = {
  id?: string;
  name?: string;
  messroom_name?: string;
  mass_room_id?: string; // used for reverse cleanup
  expires?: string;
  balance?: number | string;
  used_point?: number | string[];
  student_pin?: string;
  [k: string]: any;
};

type PayStatus = {
  [cardId: string]: {
    loading: boolean;
    error?: string;
    success?: string;
    insufficientFunds?: boolean;
  };
};

const Toast: React.FC<{ message: string; onClose: () => void }> = ({
  message,
  onClose,
}) => {
  useEffect(() => {
    const timer = setTimeout(onClose, 4000);
    return () => clearTimeout(timer);
  }, [onClose]);
  return (
    <div
      role="alert"
      className="fixed top-4 right-4 max-w-xs px-4 py-3 rounded shadow-md flex items-center gap-2"
      style={{
        backgroundColor: "#ffeeba",
        border: "1px solid #f5c26b",
        color: "#856404",
        zIndex: 9999,
      }}
    >
      <div className="text-sm">{message}</div>
      <button
        onClick={onClose}
        aria-label="close toast"
        style={{
          marginLeft: 8,
          fontWeight: "bold",
          background: "none",
          border: "none",
          cursor: "pointer",
        }}
      >
        ×
      </button>
    </div>
  );
};

const MyFoodCardTable: React.FC = () => {
  const [foodCards, setFoodCards] = useState<FoodCard[]>([]);
  const [isDeleting, setIsDeleting] = useState<string | null>(null);
  const [payStatus, setPayStatus] = useState<PayStatus>({});
  const [toast, setToast] = useState<string | null>(null);

  const dispatch = useDispatch();
  const user = useSelector((state: RootState) => state.user.user);

  const fetchFoodCards = async () => {
    if (!user?.personalIdentificationNumber) return;
    try {
      const res = await axiosInstance.get("/food/allFoodCards");
      const all: FoodCard[] = res.data?.data ?? [];
      const filtered = all.filter(
        (card) => card.student_pin === user.personalIdentificationNumber
      );
      setFoodCards(filtered);
    } catch (err) {
      console.error("Failed to fetch food cards:", err);
      setFoodCards([]);
    }
  };

  useEffect(() => {
    fetchFoodCards();
  }, [user?.personalIdentificationNumber]);

  const handleUnsubscribe = async (card: FoodCard) => {
    if (!card.id) return;
    if (!user?.personalIdentificationNumber) {
      setToast("User PIN missing, cannot unsubscribe fully.");
      return;
    }
    if (!window.confirm("Are you sure you want to unsubscribe this card?"))
      return;

    try {
      setIsDeleting(card.id);

      const pin = encodeURIComponent(user.personalIdentificationNumber);
      const messID = card.mass_room_id; // according to provided shape

      if (messID) {
        await axiosInstance.delete(
          `/food/removeMessUser/${encodeURIComponent(messID)}/${pin}`
        );
      } else {
        console.warn(
          "No mess room ID on card; skipping removal from mess_room_users"
        );
      }

      await axiosInstance.delete(`/food/deleteFoodCard/${card.id}`);

      setFoodCards((prev) => prev.filter((c) => c.id !== card.id));
      setToast("Unsubscribed and cleaned up successfully.");
    } catch (err: any) {
      console.error("Unsubscribe failed:", err);
      setToast("Failed to unsubscribe fully.");
    } finally {
      setIsDeleting(null);
    }
  };

  const handlePayForMeal = async (card: FoodCard) => {
    const cardId = card.id ?? "unknown";
    const pin = user?.personalIdentificationNumber;
    if (!pin) {
      setPayStatus((s) => ({
        ...s,
        [cardId]: { loading: false, error: "User PIN missing" },
      }));
      return;
    }

    const endpoint = `http://localhost:8000/api/food/payForMeal/${encodeURIComponent(
      pin
    )}`;

    setPayStatus((s) => ({
      ...s,
      [cardId]: {
        ...(s[cardId] || {}),
        loading: true,
        error: undefined,
        success: undefined,
      },
    }));

    try {
      const resp = await axiosInstance.post(endpoint);
      console.log("PayForMeal success:", resp.data);
      await fetchFoodCards();
      setPayStatus((s) => ({
        ...s,
        [cardId]: { loading: false, success: "Meal paid" },
      }));
    } catch (err: any) {
      console.error("PayForMeal error:", err);
      const resp = err?.response;
      const insufficient =
        resp?.status === 400 &&
        typeof resp?.data?.error === "string" &&
        resp.data.error.toLowerCase().includes("insufficient balance");

      if (insufficient) {
        setPayStatus((s) => ({
          ...s,
          [cardId]: {
            loading: false,
            insufficientFunds: true,
            error: "Insufficient balance in the food card",
          },
        }));
        setToast("Insufficient funds — please top up the account.");
        return;
      }

      setPayStatus((s) => ({
        ...s,
        [cardId]: { loading: false, error: "Payment failed" },
      }));
    }
  };

  const totalNumericBalance = useMemo(() => {
    if (foodCards.length === 0) return 0;
    const numeric = foodCards
      .map((c) => {
        if (typeof c.balance === "number") return c.balance;
        if (typeof c.balance === "string") {
          const parsed = parseFloat(c.balance.replace(",", "."));
          return isNaN(parsed) ? null : parsed;
        }
        return null;
      })
      .filter((n): n is number => n != null);
    if (numeric.length) return numeric.reduce((a, b) => a + b, 0);
    const first = foodCards[0].balance;
    if (typeof first === "number") return first;
    if (typeof first === "string") {
      const parsed = parseFloat(first.replace(",", "."));
      return isNaN(parsed) ? 0 : parsed;
    }
    return 0;
  }, [foodCards]);

  const displayBalance = useMemo(
    () => (foodCards.length === 0 ? "N/A" : totalNumericBalance.toString()),
    [totalNumericBalance, foodCards]
  );

  return (
    <div className="relative overflow-x-auto">
      {toast && <Toast message={toast} onClose={() => setToast(null)} />}

      <div className="mb-4 flex justify-between items-center">
        <div>
          <span className="text-sm font-medium text-gray-700">
            Total balance:{" "}
            <span className="font-semibold">{displayBalance}</span>
          </span>
        </div>
      </div>

      <table className="min-w-full bg-white border-collapse">
        <thead>
          <tr>
            <th className="py-2 px-4 border-b text-left text-sm font-medium">
              Card Name
            </th>
            <th className="py-2 px-4 border-b text-left text-sm font-medium">
              Expires
            </th>
            <th className="py-2 px-4 border-b text-left text-sm font-medium">
              Pay for Meal
            </th>
            <th className="py-2 px-4 border-b text-left text-sm font-medium">
              Actions
            </th>
          </tr>
        </thead>
        <tbody>
          {foodCards.length > 0 ? (
            foodCards.map((card, idx) => {
              const key = card.id ?? `no-id-${idx}`;
              const status = payStatus[card.id ?? "unknown"] || {
                loading: false,
              };
              return (
                <tr key={key} className="even:bg-gray-50">
                  <td className="py-2 px-4 border-b text-center">
                    {card.messroom_name ?? card.name ?? "N/A"}
                  </td>
                  <td className="py-2 px-4 border-b text-center">
                    {card.expires ?? "N/A"}
                  </td>
                  <td className="py-2 px-4 border-b text-center align-top">
                    <button
                      onClick={() => handlePayForMeal(card)}
                      disabled={
                        status.loading || !user?.personalIdentificationNumber
                      }
                      className="inline-flex items-center gap-1 border border-green-600 bg-green-100 hover:bg-green-200 px-3 py-1 rounded text-xs"
                    >
                      {status.loading ? "Processing..." : "Pay for Meal"}
                    </button>
                    {status.success && (
                      <div className="text-green-600 text-xs mt-1">
                        {status.success}
                      </div>
                    )}
                    {status.error && !status.insufficientFunds && (
                      <div className="text-red-600 text-xs mt-1">
                        {status.error}
                      </div>
                    )}
                    {status.insufficientFunds && (
                      <div className="text-orange-600 text-xs mt-1">
                        {status.error}
                      </div>
                    )}
                  </td>
                  <td className="py-2 px-4 text-center">
                    {card.id ? (
                      <button
                        onClick={() => handleUnsubscribe(card)}
                        disabled={isDeleting === card.id}
                        className="inline-flex items-center gap-1 border border-red-600 bg-red-100 hover:bg-red-200 px-3 py-1 rounded text-xs"
                      >
                        {isDeleting === card.id
                          ? "Unsubscribing..."
                          : "Unsubscribe"}
                      </button>
                    ) : (
                      <span className="text-xs text-yellow-600">
                        No ID to unsubscribe
                      </span>
                    )}
                  </td>
                </tr>
              );
            })
          ) : (
            <tr>
              <td
                colSpan={4}
                className="py-4 px-4 text-center text-sm text-gray-600"
              >
                No Food Cards available...
              </td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};

export default MyFoodCardTable;
