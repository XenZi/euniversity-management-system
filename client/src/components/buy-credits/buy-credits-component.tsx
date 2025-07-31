import { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { axiosInstance } from "../../services/axios.service";
import { closeModal } from "../../redux/slices/modal.slice";
import { RootState } from "../../redux/store/store";

const BuyCredits: React.FC = () => {
  const dispatch = useDispatch();
  const user = useSelector((state: RootState) => state.user.user);

  const [formData, setFormData] = useState({
    credit_card_number: "",
    name: "",
    cvv: "",
    amount: "",
  });
  const [touched, setTouched] = useState({
    credit_card_number: false,
    name: false,
    cvv: false,
    amount: false,
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  // validators
  const isCardNumberValid = /^\d{16}$/.test(
    formData.credit_card_number.replace(/\s+/g, "")
  );
  const isCVVValid = /^\d{3}$/.test(formData.cvv);
  const amountInt = parseInt(formData.amount, 10);
  const isAmountValid =
    !isNaN(amountInt) &&
    amountInt > 0 &&
    String(amountInt) === formData.amount.trim();
  const isNameValid = formData.name.trim() !== "";

  const isFormValid =
    isCardNumberValid && isCVVValid && isAmountValid && isNameValid;

  const fieldClass = (valid: boolean, wasTouched: boolean) =>
    `w-full mb-2 p-3 rounded border-2 ${
      !valid && wasTouched ? "border-red-500" : "border-gray-800"
    }`;

  const submitForm = async () => {
    setError(null);
    setSuccess(null);
    setTouched({
      credit_card_number: true,
      name: true,
      cvv: true,
      amount: true,
    });

    if (!user?.personalIdentificationNumber) {
      setError("User PIN unavailable.");
      return;
    }
    if (!isFormValid) {
      setError("Please fix validation errors before submitting.");
      return;
    }

    const buyCreditsData = {
      student_pin: user.personalIdentificationNumber,
      credit_card_number: formData.credit_card_number,
      name: formData.name,
      cvv: formData.cvv,
      amount: String(amountInt),
    };

    try {
      setLoading(true);
      await axiosInstance.post("/food/createPayment", buyCreditsData);
      setSuccess("Credits added successfully.");
      setTimeout(() => dispatch(closeModal()), 800);
    } catch (err: any) {
      console.error("Buy credits failed:", err);
      const resp = err?.response;
      if (resp?.data?.error) setError(resp.data.error);
      else setError("Failed to add credits.");
    } finally {
      setLoading(false);
    }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    if (name === "amount") {
      if (value === "" || /^\d*$/.test(value)) {
        setFormData((f) => ({ ...f, [name]: value }));
      }
    } else if (name === "credit_card_number") {
      // allow digits only, up to 16
      const cleaned = value.replace(/\D/g, "").slice(0, 16);
      setFormData((f) => ({ ...f, [name]: cleaned }));
    } else if (name === "cvv") {
      const cleaned = value.replace(/\D/g, "").slice(0, 3);
      setFormData((f) => ({ ...f, [name]: cleaned }));
    } else {
      setFormData((f) => ({ ...f, [name]: value }));
    }
  };

  const handleBlur = (field: keyof typeof touched) => {
    setTouched((t) => ({ ...t, [field]: true }));
  };

  return (
    <div className="flex flex-col items-center justify-center h-full w-full">
      <form
        className="flex flex-col w-full max-w-lg px-8 py-10 bg-white rounded shadow-lg"
        onSubmit={(e) => {
          e.preventDefault();
          submitForm();
        }}
        noValidate
      >
        <h2 className="text-center text-3xl font-semibold mb-6">Buy Credits</h2>

        {error && (
          <div className="mb-4 text-sm text-red-700 bg-red-100 p-3 rounded">
            {error}
          </div>
        )}
        {success && (
          <div className="mb-4 text-sm text-green-700 bg-green-100 p-3 rounded">
            {success}
          </div>
        )}

        <div className="flex flex-col">
          <input
            type="text"
            name="credit_card_number"
            className={fieldClass(
              isCardNumberValid,
              touched.credit_card_number
            )}
            placeholder="Credit Card Number (16 digits)"
            value={formData.credit_card_number}
            onChange={handleInputChange}
            onBlur={() => handleBlur("credit_card_number")}
            required
            disabled={loading}
          />
          {!isCardNumberValid && touched.credit_card_number && (
            <div className="text-xs text-red-600 mb-2">
              Credit card number must be exactly 16 digits.
            </div>
          )}

          <input
            type="text"
            name="name"
            className={fieldClass(isNameValid, touched.name)}
            placeholder="Name on Card"
            value={formData.name}
            onChange={handleInputChange}
            onBlur={() => handleBlur("name")}
            required
            disabled={loading}
          />
          {!isNameValid && touched.name && (
            <div className="text-xs text-red-600 mb-2">Name is required.</div>
          )}

          <input
            type="text"
            name="cvv"
            className={fieldClass(isCVVValid, touched.cvv)}
            placeholder="CVV (3 digits)"
            value={formData.cvv}
            onChange={handleInputChange}
            onBlur={() => handleBlur("cvv")}
            required
            disabled={loading}
          />
          {!isCVVValid && touched.cvv && (
            <div className="text-xs text-red-600 mb-2">
              CVV must be exactly 3 digits.
            </div>
          )}

          <input
            type="text"
            name="amount"
            inputMode="numeric"
            className={fieldClass(isAmountValid, touched.amount)}
            placeholder="Amount (integer)"
            value={formData.amount}
            onChange={handleInputChange}
            onBlur={() => handleBlur("amount")}
            required
            disabled={loading}
          />
          {!isAmountValid && touched.amount && (
            <div className="text-xs text-red-600 mb-2">
              Amount must be a positive integer.
            </div>
          )}
        </div>

        <button
          type="submit"
          disabled={loading || !isFormValid}
          className="mt-4 w-full bg-auburn-500 hover:bg-auburn-600 text-white font-semibold py-3 rounded disabled:opacity-50"
        >
          {loading ? "Processing..." : "Buy Credits"}
        </button>
      </form>
    </div>
  );
};

export default BuyCredits;
