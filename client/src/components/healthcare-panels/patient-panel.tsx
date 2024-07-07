import { Appointment, Prescription, Referral, UserRecord } from "../../models/record.model"
import { useSelector } from "react-redux";
import { RootState } from "../../redux/store/store";
import { useEffect, useState } from "react";
import { axiosInstance } from "../../services/axios.service";
import { useDispatch } from "react-redux";
import { useModalContext } from "../../context/modal.context";
import { closeModal, setModalOpen } from "../../redux/slices/modal.slice";
import CreateReferralForm from "../forms/healthcare-forms/create-referral";
import { User } from "../../models/user.model";
import CreateCertificateForm from "../forms/healthcare-forms/create-certificate";
import CreatePrescriptionForm from "../forms/healthcare-forms/create-prescription";
import { EAppointmentStatus, EAppointmentType, EDrugForm, EPrescriptionStatus } from "../../models/enum";
import ExtendPrescriptionDialog from "../dialogs/healthcare-dialogs/extend-prescription";
import CreateAppointmentForm from "../forms/healthcare-forms/create-appointment";
import FinishAppointmentForm from "../forms/healthcare-forms/finish-appointment";






const PatientPanel: React.FC<{
    userID? : string
}> = ({userID}) => {    
    const user = useSelector((state: RootState) => state.user.user);
    const [userRecord, setUserRecord] = useState<UserRecord>();
    const [doctors, setDoctors] = useState<User[]>([]);
    const [referrals, setReferrals] = useState<Referral[]>([]); 
    const [prescriptions, setPrescriptions] = useState<Prescription[]>([]);
    const [appointments, setAppointments]= useState<Appointment[]>([])


    const dispatch = useDispatch();
    const { setContent } = useModalContext();

    useEffect(() => {
        if ((user?.roles[0] == "Patient" && userID == user?.personalIdentificationNumber) || user?.roles[0]=="Doctor"){
          fetchUserData()
        }
    
      }, [userID])  


      const fetchUserData = () => {
        axiosInstance
            .get(`/healthcare/records/${userID}`)
            .then((data) => {
                setUserRecord(data.data.data);
                setReferrals(data.data.data.referrals); 
                setPrescriptions(data.data.data.prescriptions);
                setAppointments(data.data.data.appointments);
            })
            .catch((err) => {
                console.error(err);
            });

        axiosInstance
            .get("/auth/getUsers/Doctor")
            .then((data) => {
                setDoctors(data.data.data);
            })
            .catch((err) => {
                console.log(err);
            });
    };

    const handleUpdate = () => {
        fetchUserData()
    }

    const handleMakeReferral = (patientId: string | undefined) => {
        dispatch(closeModal())
        dispatch(setModalOpen());
        setContent(<CreateReferralForm patientId={patientId}/>)
    }

    const handleMakeCertificate = (patientId: string | undefined,doctorId: string| undefined) => {
        dispatch(closeModal())
        dispatch(setModalOpen());
        setContent(<CreateCertificateForm patientId={patientId} doctorID={doctorId}/>)

    }

    const handleMakeAppointment = (patientId: string | undefined) => {
        dispatch(closeModal())
        dispatch(setModalOpen());
        setContent(<CreateAppointmentForm patientId={patientId}/>)

    }

    const handleMakePrescription = (patientId: string | undefined,doctorId: string| undefined) => {
        dispatch(closeModal())
        dispatch(setModalOpen());
        setContent(<CreatePrescriptionForm patientId={patientId} doctorID={doctorId}/>)

    }

    const handleExtendPrescription = (patientId: string, prescriptionId: string) => {
        dispatch(closeModal());
        dispatch(setModalOpen());
        setContent(
          <ExtendPrescriptionDialog
            functionToProceedOnExtend={() => {
              extendPrescription(patientId, prescriptionId);
            }}
          />
        );

    }

    const handleFinishAppointment = (patientId: string | undefined,doctorId: string| undefined, appointmentId: string|undefined) => {
        dispatch(closeModal());
        dispatch(setModalOpen());
        setContent(<FinishAppointmentForm patientId={patientId} doctorID={doctorId} appointmentId={appointmentId}/>)

    }

    const extendPrescription = (patientId: string, prescriptionId: string) => {
        axiosInstance
          .post(`/healthcare/records/${patientId}/prescriptions/${prescriptionId}/ISSUED_REPEAT`)
          .then((data) => {
            console.log(data.data);
            fetchAllPrescriptions(patientId)
          })
          .catch((err) => {
            setPrescriptions([]);
            console.log(err);
          });
        dispatch(closeModal());
      }; 
      
    const fetchAllPrescriptions = (patientId: string) => {
    axiosInstance
        .get(`/healthcare/records/${patientId}/prescriptions`)
        .then((response) => {
        setPrescriptions(response.data);
        })
        .catch((error) => {
        console.error('Error fetching prescriptions:', error);
        setPrescriptions([]); 
        });
    };



    return (
        <div className="flex flex-col items-center justify-center w-full h-screen overflow-hidden">
        <div className="bg-white w-full p-4 rounded-lg shadow-md mb-4 overflow-y-auto">
        <h3 className="text-3xl text-center mb-4">Record Details</h3>
        <ul className="list-none pl-3">
            <li className="mb-2">
                Patient Name: <span className="font-bold">{user?.fullName}</span>
            </li>
            <li className="mb-2">
                Patient PIN: <span className="font-bold">{userRecord?.patientID}</span>
            </li>
            <li className="mb-2">
                Record ID: <span className="font-bold">{userRecord?.id}</span>
            </li>
        </ul>
    </div>

    <div className="bg-white w-full p-4 rounded-lg shadow-md mb-4">
        <h3 className="text-3xl text-center mb-4">Certificate Details</h3>
        <ul className="list-none pl-3">
            <li className="mb-2">
                Title: <span className="font-bold">{userRecord?.certificate?.report?.title}</span>
            </li>
            <li className="mb-2">
                Content: <span className="font-bold">{userRecord?.certificate?.report?.content}</span>
            </li>
            <li className="mb-2">
                Date of Issue: <span className="font-bold">{userRecord?.certificate?.dateOfIssue}</span>
            </li>
        </ul>
    </div>

    {referrals && referrals.length > 0 && (
        <div className="overflow-x-auto w-full">
            <h3 className="text-3xl text-center mb-4">Referral Information</h3>
            <table className="min-w-full bg-white rounded-lg shadow-md mb-4">
                <thead>
                    <tr>
                        <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                            Date of Issue
                        </th>
                        <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                            Doctor Name
                        </th>
                        <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                            Doctor ID
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {referrals.map((referral, index) => (
                        <tr key={index} className="cursor-pointer hover:bg-gray-100">
                            <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                                {referral.dateOfIssue}
                            </td>
                            <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                                {doctors.find(doc => doc.personalIdentificationNumber === referral.doctorID)?.fullName}
                            </td>
                            <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                                {referral.doctorID}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )}
    {prescriptions && prescriptions.length > 0 && (
        <div className="overflow-x-auto w-full">
            <h3 className="text-3xl text-center mb-4">Prescription Information</h3>
            <table className="min-w-full bg-white rounded-lg shadow-md mb-4">
                <thead>
                    <tr>
                        <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                            Date of Issue
                        </th>
                        <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                            Doctor Name
                        </th>
                        <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                            Drug
                        </th>
                        <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                            Drug Form
                        </th>
                        <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                            Drug Dosage
                        </th>
                        <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                            Prescription Status
                        </th>
                        {user?.roles[0] === "Doctor" && (
                             <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                             Extend Prescription
                         </th>
                        )}
                    </tr>
                </thead>
                <tbody>
                {prescriptions.map((prescription, index) => (
                    <tr key={index} className="cursor-pointer hover:bg-gray-100">
                        <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                            {prescription.dateOfIssue}
                        </td>
                        <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                            {doctors.find(doc => doc.personalIdentificationNumber === prescription.doctorID)?.fullName}
                        </td>
                        <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                            {prescription.drug}
                        </td>
                        <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                            {EDrugForm[prescription.form as keyof typeof EDrugForm]}
                        </td>
                        <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                            {prescription.dosage}
                        </td>
                        <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                            {EPrescriptionStatus[prescription.prescriptionStatus as keyof typeof EPrescriptionStatus]}
                        </td>
                        {user?.roles[0] === "Doctor" && (
                        <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                            {prescription.prescriptionStatus == "3" ? (
                                <button
                                    className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white m-2"
                                    onClick={(e) => {
                                        e.preventDefault();
                                        if (userRecord?.patientID && prescription.id) {
                                            handleExtendPrescription(userRecord.patientID, prescription.id);
                                        } else {
                                            console.error('patientID or prescription.id is undefined');
                                        }
                                    }}
                                >
                                    Extend
                                </button>
                            ) : (
                                <span className="text-gray-400">Not Extendable</span>
                            )}
                        </td>
                    )}
                        
        </tr>
    ))}
</tbody>

            </table>
        </div>
     )}
    {appointments && appointments.length > 0 && (
        <div className="overflow-x-auto w-full">
        <h3 className="text-3xl text-center mb-4">Appointment Information</h3>
        <table className="min-w-full bg-white rounded-lg shadow-md mb-4">
            <thead>
                <tr>
                    <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                        Date of Issue
                    </th>
                    <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                        Doctor Name
                    </th>
                    <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                        Appointment Type
                    </th>
                    <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                        Appointment Status
                    </th>
                    {user?.roles[0] === "Doctor" && (
                         <th className="py-2 px-4 border-b border-gray-300 text-left text-center font-medium text-gray-700">
                         Finish Appointment
                     </th>
                    )}
                </tr>
            </thead>
            <tbody>
            {appointments.map((appointments, index) => (
                <tr key={index} className="cursor-pointer hover:bg-gray-100">
                    <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        {appointments.dateOfIssue}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        {doctors.find(doc => doc.personalIdentificationNumber === appointments.doctorID)?.fullName}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        {EAppointmentType[appointments.appointmentType as keyof typeof EAppointmentType]}
                    </td>
                    <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {EAppointmentStatus[appointments.appointmentStatus as keyof typeof EAppointmentStatus]}                    
                    </td>
                    {user?.roles[0] === "Doctor" && user.personalIdentificationNumber === appointments.doctorID &&(
                    <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        {appointments.appointmentStatus == "0" ? (
                            <button
                                className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white m-2"
                                onClick={(e) => {
                                    e.preventDefault();
                                    if (userRecord?.patientID && user.personalIdentificationNumber && appointments.id) {
                                        handleFinishAppointment(userRecord.patientID, user.personalIdentificationNumber, appointments.id);
                                    } else {
                                        console.error('something is undefined');
                                    }
                                }}
                            >
                                Complete
                            </button>
                        ) : (
                            <span className="text-gray-400">Completed</span>
                        )}
                    </td>
                )}
                    
    </tr>
))}
</tbody>

        </table>
    </div>
    )}
    {user?.roles[0] === "Doctor" ? (
    <div className="flex flex-row items-center justify-center w-full">
        <button
            className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white m-2"
            onClick={(e) => {
                e.preventDefault();
                handleMakeReferral(userRecord?.patientID);
            }}
        >
            Make Referral
        </button>

        <button
            className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white m-2"
            onClick={(e) => {
                e.preventDefault();
                handleMakeCertificate(userRecord?.patientID, user.personalIdentificationNumber);
            }}
        >
            Make Certificate
        </button>

        <button
            className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white m-2"
            onClick={(e) => {
                e.preventDefault();
                handleMakePrescription(userRecord?.patientID, user.personalIdentificationNumber);
            }}
        >
            Make Prescription
        </button>
    </div>
) : (
    <div className="flex flex-row items-center justify-center w-full">
        <button
            className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white m-2"
            onClick={(e) => {
                e.preventDefault();
                handleMakeAppointment(userRecord?.patientID)
            }}
        >
            Make Appointment
        </button>
        <button
            className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white m-2"
            onClick={(e) => {
                e.preventDefault();
                handleUpdate()
            }}
        >
            Update
        </button>
    </div>
)}

</div>

    );
};

export default PatientPanel;