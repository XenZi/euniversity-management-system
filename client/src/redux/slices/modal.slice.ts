import { createSlice } from "@reduxjs/toolkit";

export interface ModalState {
  isVisible: boolean;
}

const initialState: ModalState = {
  isVisible: false,
};

const modalSlice = createSlice({
  name: "modal",
  initialState,
  reducers: {
    setModalOpen(state) {
      state.isVisible = true;
    },
    closeModal(state) {
      state.isVisible = false;
    },
  },
});

export const { setModalOpen, closeModal } = modalSlice.actions;
export default modalSlice.reducer;
