import { EAppointmentStatus, EAppointmentType, EDrugForm, EPrescriptionStatus } from "./enum";

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
    form?: EDrugForm,
    dosage?: string,
    status?: EPrescriptionStatus,
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
    doctorID?: string,
    appointmentType?: EAppointmentType,
    appointmentStatus?: EAppointmentStatus,
    report?: UserReport,
}


