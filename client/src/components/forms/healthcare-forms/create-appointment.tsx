import { useEffect, useState } from "react";
import { axiosInstance } from "../../../services/axios.service";
import { Department, Schedule} from "../../../models/record.model";
import { EAppointmentType } from "../../../models/enum";

interface AppointmentData {
    id: string,
    time: string,
    appType: string,
}

const CreateAppointmentForm: React.FC<{patientId?: string}> = ({patientId}) => {
    const [departments, setDepartments] = useState<Department[]>([])
    const [department, setDepartment] = useState<Department>()
    const [schedule, setSchedule] = useState<Schedule>()
    const [scheduleDate, setScheduleDate] = useState<string>("")
    const [slots, setSlots] = useState<string[]>([])
    const [selectedSlot, setSelectedSlot] = useState<string>('');
    const [createFormData, setFormData] = useState<AppointmentData>({
        id: patientId ?? "",
        time: "",
        appType: "",
    })

    useEffect(() => {
        setDepartments([])
        axiosInstance
        .get("/healthcare/departments")
        .then((data) => {
            setDepartments(data.data.data)
        })
        .catch((err) => {
            console.log(err)
        })
    }, [])

    useEffect(() => {
        if (department?.schedule) {
            setSchedule(department?.schedule)
        }
    }, [department])

    useEffect(() => {
        axiosInstance
            .get(`healthcare/department/${department?.name}/schedule/${scheduleDate}/free`)
            .then((response) => {
                const responseData = response.data.data;
                setSlots(responseData || []);
                console.log(responseData);
            })
            .catch((error) => {
                console.error('Error fetching slots:', error);
                setSlots([]); // Set empty array if there's an error
            });
    }, [scheduleDate]);
    
    
    const handleDepartmentChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedDepartmentName = event.target.value;
        const selectedDepartment = departments.find(dep => dep.name === selectedDepartmentName);
        setDepartment(selectedDepartment);
    };

    const handleScheduleDateChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedDate = event.target.value;
        setScheduleDate(selectedDate);
    };

    const handleSlotSelectionChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedSlotTime = event.target.value;
        console.log("Selected slot time:", selectedSlotTime);    
        setFormData({ ...createFormData, time: selectedSlotTime });
        setSelectedSlot(selectedSlotTime)
        
    };

    const handleTypeChange = (
        event: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
    ) => {
        var { name, value } = event.target;
        setFormData({
            ...createFormData,
            [name]: value,
        });
    };


    const handleClearSelection = () => {
        setDepartment(undefined);
        setSchedule(undefined);
        setScheduleDate("");
        setSlots([]);
        setSelectedSlot('');
        setFormData({
            ...createFormData,
            time: '',
            appType: '',
        });
    };
        
    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        console.log(createFormData)

        event.preventDefault();
        axiosInstance
            .post(`/healthcare/department/${department?.name}/schedule/${scheduleDate}/add`, createFormData)
            .then((resp) => {
                console.log(resp.data.data)
            })
            .catch((err) => {
                console.log("PUKO")
                console.log(err)
            })
    };

    return (
        <div>
        <form
            action="#"
            className="max-w-lg mx-auto p-6 bg-white shadow-md rounded-lg"
            onSubmit={handleSubmit}
        >
            <h2 className="text-center text-3xl font-semibold mb-6">
                Create Appointment
            </h2>
    
            <div className="mb-6">
                <label htmlFor="department" className="block text-sm font-semibold mb-1">
                    Department
                </label>
                <select
                    id="department"
                    name="department"
                    onChange={handleDepartmentChange}
                    value={department?.name || ""}
                    className="w-full p-3 border-2 border-battleship-500 rounded"
                    required
                >
                    <option value="">Select Department</option>
                    {departments.map(dep => (
                        <option key={dep.name} value={dep.name}>
                            {dep.name}
                        </option>
                    ))}
                </select>
            </div>
    
            {department && (
                <div className="mb-6">
                    <label htmlFor="scheduleDate" className="block text-sm font-semibold mb-1">
                        Schedule Date
                    </label>
                    <select
                        id="scheduleDate"
                        name="scheduleDate"
                        onChange={handleScheduleDateChange}
                        value={scheduleDate}
                        className="w-full p-3 border-2 border-battleship-500 rounded"
                        required
                    >
                        <option value="">Select Date</option>
                        {Object.keys(schedule?.date ?? {}).map(date => (
                            <option key={date} value={date}>
                                {date}
                            </option>
                        ))}
                    </select>
                </div>
            )}
    
            {scheduleDate && slots.length > 0 && (
                <div className="mb-6">
                    <label htmlFor="scheduleDate" className="block text-sm font-semibold mb-1">
                        Available Slots:
                    </label>                        
                            <select
                            id="slot"
                            name="slot"
                            onChange={handleSlotSelectionChange}
                            value={createFormData.time}
                            className="w-full p-3 border-2 border-battleship-500 rounded"
                            required

                            >
                            <option value="">Select Slot</option>
                                {slots.map((slot, index) => (
                                    <option key={index} value={slot}>
                                        {slot}
                                    </option>
                                ))}
                        </select>
                    </div>
                )}
            {selectedSlot && (
                <div className="mb-6">
                <label htmlFor="form" className="block text-sm font-semibold mb-1">
                    Prescription Status
                </label>
                <select
                id="appType"
                name="appType"
                value={createFormData.appType}
                onChange={handleTypeChange}
                className="w-full p-3 border-2 border-battleship-500 rounded"
                required
                >
                    {Object.entries(EAppointmentType)
                .filter(([key]) => isNaN(Number(key))) 
                .map(([key]) => (
                <option key={key} value={key}>
                    {key}
                        </option>
                    ))}
                    </select>
                </div>
                )}    


<div className="flex justify-between">
                    <button
                        type="submit"
                        className="w-full border bg-auburn-500 border-auburn-500 font-semibold py-2 px-4 rounded focus:border-auburn-700 text-white hover:bg-auburn-600 transition duration-300"
                    >
                        Create Appointment
                    </button>
                    <button
                        type="button"
                        onClick={handleClearSelection}
                        className="w-full border bg-gray-300 border-gray-300 font-semibold py-2 px-4 rounded text-gray-600 hover:bg-gray-400 transition duration-300"
                    >
                        Clear Selection
                    </button>
                </div>
            </form>
        </div>
    );
};


export default CreateAppointmentForm;