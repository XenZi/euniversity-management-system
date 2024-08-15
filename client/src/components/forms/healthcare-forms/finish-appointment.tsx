import { useState } from "react";
import { axiosInstance } from "../../../services/axios.service";

interface FinishAppointmentData {
    AppointmentId: string,
    title: string,
    content: string,
    patientID: string,
    doctorID: string,
   
}

const FinishAppointmentForm: React.FC<{patientId?: string, doctorID?: string, appointmentId?: string}> = ({patientId, doctorID, appointmentId}) => {
    const [createFormData, setFormData] = useState<FinishAppointmentData>({
        patientID: patientId ?? "",
        doctorID: doctorID ?? "",
        AppointmentId: appointmentId ?? "",
        title: "",
        content: "",
        
    })

    const handleChange = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = event.target;
        setFormData({
            ...createFormData,
            [name]: value,
        });
    };


    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        console.log(`healthcare/records/{{exampleID}}/appointments/668120a1a0eba0662e067f22/update`, createFormData);
        axiosInstance
            .post(`healthcare/records/${patientId}/appointments/${appointmentId}/update`, createFormData)
            .then((resp) => {
                console.log(resp.data.data)
            })
            .catch((err) => {
                console.log("PUKO")
                console.log(err)
            })

    };

    return (
        <form action="#" className="max-w-lg mx-auto p-6 bg-white shadow-md rounded-lg" onSubmit={handleSubmit}>
        <h2 className="text-center text-3xl font-semibold mb-6">Appointemnt Report</h2>
        
        <div className="mb-6">
            <label htmlFor="title" className="block text-sm font-semibold mb-1">
                Report Title
            </label>
            <input
                type="text"
                id="title"
                name="title"
                value={createFormData.title}
                onChange={handleChange}
                className="w-full p-3 border-2 border-battleship-500 rounded"
                placeholder="Enter certificate title..."
                required
            />
        </div>
        
        <div className="mb-6">
            <label htmlFor="content" className="block text-sm font-semibold mb-1 text-center">
                Report Content
            </label>
            <textarea
                id="content"
                name="content"
                value={createFormData.content}
                onChange={handleChange}
                className="w-full p-3 border-2 border-battleship-500 rounded"
                placeholder="Enter certificate content..."
                rows={5}
                required
            />
        </div>
    
        <button
            type="submit"
            className="w-full border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white hover:bg-auburn-600 transition duration-300"
        >
            Complete Appointment
        </button>
    </form>
    
      );


}

export default FinishAppointmentForm;