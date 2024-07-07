export interface UserRecord {
    id?: string,
    patientID?: string,
    certificate?: UserCertificate,
    prescriptions?: Prescription[],
    referrals?: Referral[],
    appointment?: Appointment[],
}

export interface UserCertificate {
    id?: string,
    dateOfIssue?: string,
    patientID?: string,
    doctorID?: string,
    report?: UserReport,
}

export interface UserReport {
    id?: string,
    title?: string,
    content?: string,
    dateOfIssue?: string,
}

export interface Prescription{
    id?: string,
    dateOfIssue?: string,
    patientID?: string,
    doctorID?: string,
    drug?: string,
    form: string,
    dosage?: string,
    prescriptionStatus: string,
}

export interface Referral {
    id?: string,
    dateOfIssue?: string,
    patientID?: string,
    doctorID?: string,
}

export interface Appointment {
    id?: string,
    dateOfIssue?: string,
    patientID?: string,
    doctorID: string,
    appointmentType: string,
    appointmentStatus: string,
    report?: UserReport,
}

export interface Department{
    id?: string,
    name: string,
    schedule: Schedule
}

export interface Schedule {
    date: { [date: string]: Slot[] };
}

export interface Slot {
    time: string,
    doctorID?: string,
    patientID?: string,
}

