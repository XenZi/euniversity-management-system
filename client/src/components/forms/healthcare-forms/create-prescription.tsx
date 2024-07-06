
import { useState } from "react";
import { axiosInstance } from "../../../services/axios.service";
import { Prescription } from "../../../models/record.model";
import { EDrugForm, EPrescriptionStatus } from "../../../models/enum";


const CreatePrescriptionForm: React.FC<{ patientId?: string; doctorID?: string }> = ({
    patientId,
    doctorID,
}) => {
    const [createFormData, setFormData] = useState<Prescription>({
        patientID: patientId ?? "",
        doctorID: doctorID ?? "",
        drug: "",
        form: "",
        dosage: "",
        prescriptionStatus: "",
    });

    const handleChange = (
        event: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
    ) => {
        var { name, value } = event.target;
        setFormData({
            ...createFormData,
            [name]: value,
        });
    };

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        console.log(`healthcare/records/${patientId}/prescriptions/createPrescription:`, createFormData);
        axiosInstance
            .post(`healthcare/records/${patientId}/prescriptions/createPrescription`, createFormData)
            .then((resp) => {
                console.log(resp.data.data);
            })
            .catch((err) => {
                console.log("Error:", err);
            });
    };

    return (
        <form
            action="#"
            className="max-w-lg mx-auto p-6 bg-white shadow-md rounded-lg"
            onSubmit={handleSubmit}
        >
            <h2 className="text-center text-3xl font-semibold mb-6">
                Create New Prescription
            </h2>

            <div className="mb-6">
                <label htmlFor="drug" className="block text-sm font-semibold mb-1">
                    Drug
                </label>
                <input
                    type="text"
                    id="drug"
                    name="drug"
                    value={createFormData.drug}
                    onChange={handleChange}
                    className="w-full p-3 border-2 border-battleship-500 rounded"
                    placeholder="Enter drug..."
                    required
                />
            </div>

            <div className="mb-6">
                <label htmlFor="form" className="block text-sm font-semibold mb-1">
                    Drug Form
                </label>
                <select
                    id="form"
                    name="form"
                    value={createFormData.form}
                    onChange={handleChange}
                    className="w-full p-3 border-2 border-battleship-500 rounded"
                    required
                    >
                     {Object.entries(EDrugForm)
          .filter(([key]) => isNaN(Number(key))) 
          .map(([key]) => (
            <option key={key} value={key}>
              {key}
            </option>
          ))}
                </select>
            </div>

            <div className="mb-6">
                <label htmlFor="dosage" className="block text-sm font-semibold mb-1">
                    Dosage
                </label>
                <input
                    type="text"
                    id="dosage"
                    name="dosage"
                    value={createFormData.dosage}
                    onChange={handleChange}
                    className="w-full p-3 border-2 border-battleship-500 rounded"
                    placeholder="Enter dosage..."
                    required
                />
            </div>

            <div className="mb-6">
                <label htmlFor="form" className="block text-sm font-semibold mb-1">
                    Prescription Status
                </label>
                <select
                    id="prescriptionStatus"
                    name="prescriptionStatus"
                    value={createFormData.prescriptionStatus}
                    onChange={handleChange}
                    className="w-full p-3 border-2 border-battleship-500 rounded"
                    required
                >
                    {Object.entries(EPrescriptionStatus)
          .filter(([key]) => isNaN(Number(key))) 
          .map(([key]) => (
            <option key={key} value={key}>
              {key}
                    </option>
                ))}
                </select>
            </div>

            <button
                type="submit"
                className="w-full border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white hover:bg-auburn-600 transition duration-300"
            >
                Create Prescription
            </button>
        </form>
    );
};

export default CreatePrescriptionForm;
