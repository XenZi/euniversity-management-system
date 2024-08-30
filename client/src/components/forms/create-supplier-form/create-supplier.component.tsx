import { useState, useEffect } from "react";
import { axiosInstance } from "../../../services/axios.service";
import { useDispatch } from "react-redux";
import { closeModal } from "../../../redux/slices/modal.slice";

interface CreateSupplierFormData {
  name: string;
  location: string;
  phone_number: string;
  messRoom: Mess | null; // Add messRoom field
}

interface Mess {
  id: string;
  name: string;
  location: string;
  capacity: number;
  rating: number;
  supplier_id: string | null;
  mess_room_users: string[]; // Replace `any[]` with the actual type if known
}

const CreateSupplierForm = () => {
  const [createSupplierFormData, setCreateSupplierFormData] =
    useState<CreateSupplierFormData>({
      name: "",
      location: "",
      phone_number: "",
      messRoom: null, // Initialize messRoom field
    });

  const [messRooms, setMessRooms] = useState<Mess[]>([]); // State to hold messRooms

  const dispatch = useDispatch();

  // Fetch messRooms when component mounts
  useEffect(() => {
    const fetchMessRooms = async () => {
      try {
        const response = await axiosInstance.get("/food/allMessRooms");
        setMessRooms(response.data.data);
      } catch (error) {
        console.error("Failed to fetch messRooms:", error);
      }
    };

    fetchMessRooms();
  }, []);

  const onInputChange = (
    e: React.FormEvent<HTMLInputElement | HTMLSelectElement>,
    key: keyof CreateSupplierFormData
  ) => {
    const value = e.currentTarget.value;
    setCreateSupplierFormData((prevState) => ({
      ...prevState,
      [key]: value,
    }));
  };

  const submitForm = async () => {
    try {
      const { name, location, phone_number } = createSupplierFormData;

      // Create an object with the specific fields
      const supplierData = {
        name,
        location,
        phone_number,
      };
      // Step 1: Create the supplier
      const supplierResponse = await axiosInstance.post(
        "/food/createSupplier",
        supplierData
      );
      console.log(supplierResponse);
      const newSupplierId = supplierResponse.data.data.id;

      // Step 2: Update the supplier_id in the mess
      if (createSupplierFormData.messRoom) {
        await axiosInstance.put(
          `/food/updateMessSupplier/${createSupplierFormData.messRoom.id}/${newSupplierId}`
        );
      } else {
        console.error("No mess room selected.");
        // Optionally, handle the case where no mess room is selected
      }

      // Close the modal after successful operations
      dispatch(closeModal());
    } catch (err) {
      console.log(err);
    }
  };

  // Filter mess rooms where supplier_id is an empty string
  const availableMessRooms = messRooms.filter(
    (room) => room.supplier_id === ""
  );

  return (
    <form action="#" className="flex flex-col">
      <h2 className="text-center text-3xl font-semibold mb-3">
        Create New Supplier
      </h2>
      <input
        type="text"
        name="name"
        id="name"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Supplier name..."
        value={createSupplierFormData.name}
        onChange={(e) => onInputChange(e, "name")}
      />
      <input
        type="text"
        name="location"
        id="location"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Supplier location..."
        value={createSupplierFormData.location}
        onChange={(e) => onInputChange(e, "location")}
      />
      <input
        type="text"
        name="phoneNumber"
        id="phoneNumber"
        className="mb-3 p-3 border-2 border-battleship-500"
        placeholder="Phone number..."
        value={createSupplierFormData.phone_number}
        onChange={(e) => onInputChange(e, "phone_number")}
      />
      <select
        name="messRoom"
        id="messRoom"
        className="mb-3 p-3 border-2 border-battleship-500"
        value={createSupplierFormData.messRoom?.id || ""}
        onChange={(e) => {
          const selectedRoom = availableMessRooms.find(
            (room) => room.id === e.target.value
          );
          setCreateSupplierFormData((prevState) => ({
            ...prevState,
            messRoom: selectedRoom || null,
          }));
        }}
      >
        <option value="">Select Mess Room...</option>
        {availableMessRooms.map((room) => (
          <option key={room.id} value={room.id}>
            {room.name}
          </option>
        ))}
      </select>
      <button
        className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
        onClick={(e) => {
          e.preventDefault();
          submitForm();
        }}
      >
        Create Supplier
      </button>
    </form>
  );
};

export default CreateSupplierForm;
