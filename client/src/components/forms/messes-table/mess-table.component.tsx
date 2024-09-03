import { useEffect, useState } from "react";
import { Mess } from "../../../models/mess.model";
import { useDispatch } from "react-redux";
import { useModalContext } from "../../../context/modal.context";
import { axiosInstance } from "../../../services/axios.service";
import { closeModal, setModalOpen } from "../../../redux/slices/modal.slice";
import EditMessForm from "../edit-mess/edit-mess.form";
import DeleteDialog from "../../dialogs/delete-dialog/delete-dialog.component";

const MessTable: React.FC<{ adminView?: boolean }> = ({ adminView }) => {
  const [messes, setMesses] = useState<Mess[]>([]);
  const dispatch = useDispatch();
  const { setContent } = useModalContext();

  useEffect(() => {
    axiosInstance
      .get("/food/allMessRooms")
      .then((data) => {
        setMesses(data.data.data);
      })
      .catch((err) => console.log(err));
  }, []);

  const openDialogForEdit = (mess: Mess) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    setContent(<EditMessForm mess={mess} />);
  };

  const openDialogForDelete = (messID: string) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    setContent(
      <DeleteDialog
        functionToProceedOnDelete={() => {
          deleteMess(messID);
        }}
      />
    );
  };

  const deleteMess = (messID: string) => {
    axiosInstance
      .delete(`/food/deleteMessRoom/${messID}`)
      .then((data) => {
        console.log(data);
        setMesses(messes.filter((mess) => mess.id !== messID));
      })
      .catch((err) => {
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
                Capacity
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Rating
              </th>
              {adminView && (
                <>
                  <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                    Edit
                  </th>
                  <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                    Delete
                  </th>
                </>
              )}
            </tr>
          </thead>
          <tbody>
            {messes && messes.length > 0 ? (
              messes.map((mess, index) => (
                <tr key={index}>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {mess.name}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {mess.location}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {mess.capacity}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {mess.rating}
                  </td>
                  {adminView && (
                    <>
                      <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        <button
                          className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                          onClick={(e) => {
                            e.preventDefault();
                            openDialogForEdit(mess);
                          }}
                        >
                          Edit Mess
                        </button>
                      </td>
                      <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        <button
                          className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                          onClick={(e) => {
                            e.preventDefault();
                            openDialogForDelete(mess.id);
                          }}
                        >
                          Delete Mess
                        </button>
                      </td>
                    </>
                  )}
                </tr>
              ))
            ) : (
              <p className="my-2">Messes are not available..</p>
            )}
          </tbody>
        </table>
      </div>
    </>
  );
};

export default MessTable;
