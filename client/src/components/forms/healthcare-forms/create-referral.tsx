import { useEffect, useState } from "react";
import { axiosInstance } from "../../../services/axios.service";
import { User } from "../../../models/user.model";

interface CreateReferralData {
    patientID: string
    doctorID: string
}

const CreateReferralForm: React.FC<{patientId?: string}> = ({patientId}) => {
    const [doctors, setDoctors] = useState<User[]>([]);
    const [createFormData, setFormData] = useState<CreateReferralData>({
        patientID: patientId ?? "",
        doctorID: "",
    })

    const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedDoctorID = event.target.value;
        setFormData({
            ...createFormData,
            doctorID: selectedDoctorID,
        });
    };



    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        console.log(`healthcare/records/${patientId}/referrals/createReferral:`, createFormData);
        axiosInstance
            .post(`healthcare/records/${patientId}/referrals/createReferral`, createFormData)
            .then((resp) => {
                console.log(resp.data.data)
            })
            .catch((err) => {
                console.log("PUKO")
                console.log(err)
            })

    };

    useEffect(() => {
        axiosInstance
        .get("/auth/getUsers/Doctor")
        .then((data) => {
            setDoctors(data.data.data)
            console.log(doctors)
        })
        .catch((err) => {
            console.log(err)
        })
    }, [])

    return (
        <form action="#" className="flex flex-col max-w-lg mx-auto" onSubmit={handleSubmit}>
          <h2 className="text-center text-3xl font-semibold mb-3">Create new referral</h2>
          <div className="mb-3">
            <label htmlFor="doctorID" className="block text-sm font-medium text-gray-700">
              Select Doctor
            </label>
            <select
              id="doctorID"
              name="doctorID"
              className="mt-1 block w-full p-3 border-2 border-battleship-500 rounded"
              value={createFormData.doctorID}
              onChange={handleChange}
            >
              <option value="">Select a doctor</option>
              {doctors.map((doctor) => (
                <option key={doctor.personalIdentificationNumber} value={doctor.personalIdentificationNumber}>
                  {doctor.fullName}
                </option>
              ))}
            </select>
          </div>
          <button
            type="submit"
            className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white"
          >
            Create Referral
          </button>
        </form>
      );


}

export default CreateReferralForm;