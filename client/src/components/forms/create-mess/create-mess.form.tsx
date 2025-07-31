import { useState } from "react";
import { axiosInstance } from "../../../services/axios.service";
import { useDispatch } from "react-redux";
import { closeModal } from "../../../redux/slices/modal.slice";

interface CreateMessFormData {
  name: string;
  location: string;
  capacity: number;
}

const CreateMessForm: React.FC = () => {
  const [createMessFormData, setCreateMessFormData] =
    useState<CreateMessFormData>({
      name: "",
      location: "",
      capacity: 0,
    });
  const [touched, setTouched] = useState({
    name: false,
    location: false,
    capacity: false,
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const dispatch = useDispatch();

  const onInputChange = (
    e: React.FormEvent<HTMLInputElement>,
    key: keyof CreateMessFormData
  ) => {
    const value =
      key === "capacity"
        ? Math.max(0, parseInt(e.currentTarget.value) || 0)
        : e.currentTarget.value;
    setCreateMessFormData((prevState) => ({
      ...prevState,
      [key]: value,
    }));
  };

  const validateName = createMessFormData.name.trim() !== "";
  const validateLocation = createMessFormData.location.trim() !== "";
  const validateCapacity =
    Number.isInteger(createMessFormData.capacity) &&
    createMessFormData.capacity > 0;

  const isFormValid = validateName && validateLocation && validateCapacity;

  const submitForm = async () => {
    setError(null);
    setSuccess(null);
    setTouched({ name: true, location: true, capacity: true });

    if (!isFormValid) {
      setError("Please fix validation errors.");
      return;
    }

    try {
      setLoading(true);
      const resp = await axiosInstance.post(
        "/food/createMessRoom",
        createMessFormData
      );
      setSuccess("Mess created successfully.");
      setTimeout(() => {
        dispatch(closeModal());
      }, 600);
    } catch (err: any) {
      console.error("Create mess failed:", err);
      setError("Failed to create mess.");
    } finally {
      setLoading(false);
    }
  };

  const handleBlur = (field: keyof typeof touched) => {
    setTouched((t) => ({ ...t, [field]: true }));
  };

  const inputClass = (valid: boolean, wasTouched: boolean) =>
    `w-full mb-3 p-3 border-2 rounded transition ${
      !valid && wasTouched ? "border-red-500" : "border-gray-300"
    }`;

  return (
    <div className="flex flex-col items-center justify-center h-full w-full">
      <form
        className="w-full max-w-md px-8 py-8 bg-white rounded shadow-lg"
        onSubmit={(e) => {
          e.preventDefault();
          submitForm();
        }}
        noValidate
      >
        <h2 className="text-center text-2xl font-semibold mb-5">
          Create new Mess
        </h2>

        {error && (
          <div className="mb-3 text-sm text-red-700 bg-red-100 p-2 rounded">
            {error}
          </div>
        )}
        {success && (
          <div className="mb-3 text-sm text-green-700 bg-green-100 p-2 rounded">
            {success}
          </div>
        )}

        <div className="flex flex-col">
          <label htmlFor="name" className="text-sm font-medium mb-1">
            Name
          </label>
          <input
            type="text"
            name="name"
            id="name"
            className={inputClass(validateName, touched.name)}
            placeholder="Mess name..."
            value={createMessFormData.name}
            onChange={(e) => onInputChange(e, "name")}
            onBlur={() => handleBlur("name")}
            disabled={loading}
          />
          {!validateName && touched.name && (
            <div className="text-xs text-red-600 mb-2">Name is required.</div>
          )}

          <label htmlFor="location" className="text-sm font-medium mb-1">
            Location
          </label>
          <input
            type="text"
            name="location"
            id="location"
            className={inputClass(validateLocation, touched.location)}
            placeholder="Mess location..."
            value={createMessFormData.location}
            onChange={(e) => onInputChange(e, "location")}
            onBlur={() => handleBlur("location")}
            disabled={loading}
          />
          {!validateLocation && touched.location && (
            <div className="text-xs text-red-600 mb-2">
              Location is required.
            </div>
          )}

          <label htmlFor="capacity" className="text-sm font-medium mb-1">
            Capacity
          </label>
          <input
            type="number"
            name="capacity"
            id="capacity"
            className={inputClass(validateCapacity, touched.capacity)}
            placeholder="Capacity..."
            value={createMessFormData.capacity}
            onChange={(e) => onInputChange(e, "capacity")}
            onBlur={() => handleBlur("capacity")}
            disabled={loading}
            min={1}
          />
          {!validateCapacity && touched.capacity && (
            <div className="text-xs text-red-600 mb-2">
              Capacity must be a positive integer.
            </div>
          )}
        </div>

        <button
          type="submit"
          disabled={loading || !isFormValid}
          className="mt-4 w-full bg-auburn-500 hover:bg-auburn-600 text-white font-semibold py-2 px-4 rounded disabled:opacity-50"
        >
          {loading ? "Creating..." : "Create Mess"}
        </button>
      </form>
    </div>
  );
};

export default CreateMessForm;
