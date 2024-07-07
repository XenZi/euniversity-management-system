import React from "react";
import { useDispatch } from "react-redux";
import { closeModal } from "../../../redux/slices/modal.slice";

const ExtendPrescriptionDialog: React.FC<{
  functionToProceedOnExtend: () => void;
}> = ({ functionToProceedOnExtend }) => {
  const dispatch = useDispatch();
  return (
    <div className="flex items-center justify-center bg-gray-800 bg-opacity-75">
      <div className="bg-white rounded-lg overflow-hidden w-1/3">
        <div className="px-6 py-4">
          <h3 className="text-lg text-gray-900 text-center font-bold">
            Extend Confirmation
          </h3>
          <p className="mt-2 text-sm text-gray-600 text-center">
            Are you sure you want to extend the prescription?
          </p>
        </div>
        <div className="px-6 py-4 bg-gray-100 flex items-center justify-center space-x-3">
          <button
            className="border  border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-border-auburn-500"
            onClick={(e) => {
              e.preventDefault();
              dispatch(closeModal());
            }}
          >
            Cancel
          </button>
          <button
            className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
            onClick={(e) => {
              e.preventDefault();
              functionToProceedOnExtend();
            }}
          >
            Extend
          </button>
        </div>
      </div>
    </div>
  );
};

export default ExtendPrescriptionDialog;
