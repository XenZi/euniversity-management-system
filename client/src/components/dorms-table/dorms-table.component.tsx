import { useEffect, useState } from "react";
import { Dorm } from "../../models/dorm.model";
import { axiosInstance } from "../../services/axios.service";
import { useDispatch } from "react-redux";
import { closeModal, setModalOpen } from "../../redux/slices/modal.slice";
import DeleteDialog from "../dialogs/delete-dialog/delete-dialog.component";
import { useModalContext } from "../../context/modal.context";
import EditDorm from "../forms/edit-dorm/edit-dorm.form";

const DormTable: React.FC<{
  adminView?: boolean;
}> = ({ adminView }) => {
  const [dorms, setDorms] = useState<Dorm[]>([]);
  const dispatch = useDispatch();
  const { setContent } = useModalContext();

  useEffect(() => {
    axiosInstance
      .get("/dorm/all")
      .then((data) => {
        setDorms(data.data.data);
      })
      .catch((err) => console.log(err));
  }, []);

  const openDialogForEdit = (dorm: Dorm) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    setContent(<EditDorm dorm={dorm} />);
  };

  const openDialogForDelete = (dormID: string) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    setContent(
      <DeleteDialog
        functionToProceedOnDelete={() => {
          deleteDorm(dormID);
        }}
      />
    );
  };
  const deleteDorm = (dormID: string) => {
    axiosInstance
      .delete(`/dorm/${dormID}`)
      .then((data) => {
        console.log(data);
      })
      .catch((err) => {
        setDorms([]);
        console.log(err);
      });
    dispatch(closeModal());
  };

  return (
    <>
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white">
          <thead>
            <tr>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Name
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Location
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Budget
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Self financing
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Disability
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Sensitive groups
              </th>
            </tr>
          </thead>
          <tbody>
            {dorms && dorms.length > 0 ? (
              dorms.map((dorm, index) => (
                <tr key={index}>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {dorm.name}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {dorm.location}
                  </td>
                  {dorm.prices.map((price, idx) => (
                    <td
                      className="py-2 px-4 border-b border-gray-300 text-sm text-center"
                      key={idx}
                    >
                      {price.price}
                    </td>
                  ))}
                  {adminView == true ? (
                    <>
                      <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        <button
                          className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                          onClick={(e) => {
                            e.preventDefault();
                            openDialogForEdit(dorm);
                          }}
                        >
                          Edit dorm
                        </button>
                      </td>
                      <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        <button
                          className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                          onClick={(e) => {
                            e.preventDefault();
                            openDialogForDelete(dorm.id);
                          }}
                        >
                          Delete dorm
                        </button>
                      </td>
                    </>
                  ) : (
                    ""
                  )}
                </tr>
              ))
            ) : (
              <p className="my-2">Dorms are not available..</p>
            )}
          </tbody>
        </table>
      </div>
    </>
  );
};

export default DormTable;
