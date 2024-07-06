import { useEffect, useState } from "react"
import { ScholarshipApplication } from "../../../models/scholarship-application.model"
import { User } from "../../../models/user.model";
import { axiosInstance } from "../../../services/axios.service";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const CreateScholarship = () => {
    const [loadedScholarships, setLoadedScholarships] = useState<ScholarshipApplication[]>([]);
    const [selectedUser, setSelectedUser] = useState<User | null>(null);
    const [formVisible, setFormVisible] = useState(true);


    useEffect(() =>{
        axiosInstance
            .get("/university/scholarshipApplication")
            .then((res) => {
                console.log(res)
                setLoadedScholarships(res.data.data)
            })
            .catch((err)=>{
                console.log(err)
            })
    },[])

    const handleSubmit = (e: React.FormEvent) =>{
        e.preventDefault();

        if(selectedUser){
            axiosInstance
            .post("/university/scholarship",{
                student: selectedUser
            })
            .then((res) =>{
                console.log("Submission successful",res);
                toast.success('Scholarship approved!');
                setFormVisible(false);
            })
            .catch((err) =>{
                console.log("Submission error",err);
                toast.error('Something went wrong!');
            });
        }
    };

    const handleUserChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedId = e.target.value;
        const exam = loadedScholarships.find((scholarship) => scholarship.id === selectedId);
        setSelectedUser(exam?.student || null);
    };

    if (!formVisible) {
        return <div>Form submitted successfully!</div>;
      }

    return(
        <>
        <div>Create Scholarship From</div>
        <form className="w-full" onSubmit={handleSubmit}>
            <label>
                Choose
            </label>
            <select id="select-student-scholarship" className="mb-3 p-3 border-2 border-battleship-500 w-full"
                onChange={handleUserChange}>
                <option value="" disabled>
                    Select a student
                    </option>  
                {loadedScholarships?.map((scholarship) =>(
                    <option value={scholarship.id} key={scholarship.id}>
                        {scholarship.student.fullName}
                    </option>
                ))}     
            </select>
            <button
          className="border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white w-full"
          type="submit"
        >
          Add Scholarship
        </button>
        </form>
        </>
    )


}
export default CreateScholarship