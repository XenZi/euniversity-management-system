import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useModalContext } from "../../context/modal.context";
import { axiosInstance } from "../../services/axios.service";
import { closeModal, setModalOpen } from "../../redux/slices/modal.slice";
import EditMessForm from "../forms/edit-mess/edit-mess.form";
import DeleteDialog from "../dialogs/delete-dialog/delete-dialog.component";
import RatingForm from "../forms/mess-rating/mess-rating.form";
import { RootState } from "../../redux/store/store";
import { Mess as BaseMess } from "../../models/mess.model";

// Extend the shared Mess to include rating array if not already present
type MessWithRatings = BaseMess & {
  rating?: number[]; // array of individual ratings
};

const computeAverage = (ratings: number[] = []): string => {
  if (ratings.length === 0) return "—";
  const sum = ratings.reduce((a, b) => a + b, 0);
  const avg = sum / ratings.length;
  return avg.toFixed(1);
};

const MessTableRate: React.FC<{ adminView?: boolean }> = ({ adminView }) => {
  const [messes, setMesses] = useState<MessWithRatings[]>([]);
  const dispatch = useDispatch();
  const { setContent } = useModalContext();
  const user = useSelector((state: RootState) => state.user.user);

  useEffect(() => {
    axiosInstance
      .get("/food/allMessRooms")
      .then((data) => {
        setMesses(data.data.data);
      })
      .catch((err) => console.log(err));
  }, []);

  const openDialogForEdit = (mess: MessWithRatings) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    // ensure capacity is defined to satisfy the expected type
    const normalized: BaseMess = {
      ...mess,
      capacity: mess.capacity ?? 0,
    } as BaseMess;
    setContent(<EditMessForm mess={normalized} />);
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

  const openDialogForRate = (mess: MessWithRatings) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    const normalized: BaseMess = {
      ...mess,
      capacity: mess.capacity ?? 0,
    } as BaseMess;
    setContent(<RatingForm mess={normalized} />);
  };

  const deleteMess = (messID: string) => {
    axiosInstance
      .delete(`/food/deleteMessRoom/${messID}`)
      .then(() => {
        setMesses((prev) => prev.filter((mess) => mess.id !== messID));
      })
      .catch((err) => {
        console.log(err);
      });
    dispatch(closeModal());
  };

  return (
    <div className="overflow-x-auto">
      <table className="min-w-full bg-white border-collapse">
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
            {user?.roles?.includes("Student") && (
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Rate
              </th>
            )}
          </tr>
        </thead>
        <tbody>
          {messes && messes.length > 0 ? (
            messes.map((mess, index) => {
              const avg = computeAverage(mess.rating || []);
              return (
                <tr
                  key={mess.id || index}
                  className="even:bg-gray-50 hover:bg-gray-100"
                >
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {mess.name}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {mess.location}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                    {mess.capacity ?? "—"}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm text-center font-semibold">
                    {avg}
                  </td>
                  {adminView && (
                    <>
                      <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        <button
                          className="border bg-auburn-500 font-semibold py-1 px-2 rounded text-white text-xs"
                          onClick={(e) => {
                            e.preventDefault();
                            openDialogForEdit(mess);
                          }}
                        >
                          Edit
                        </button>
                      </td>
                      <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                        <button
                          className="border bg-auburn-500 font-semibold py-1 px-2 rounded text-white text-xs"
                          onClick={(e) => {
                            e.preventDefault();
                            openDialogForDelete(mess.id);
                          }}
                        >
                          Delete
                        </button>
                      </td>
                    </>
                  )}
                  {user?.roles?.includes("Student") && (
                    <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                      <button
                        className="border bg-auburn-500 font-semibold py-1 px-2 rounded text-white text-xs"
                        onClick={(e) => {
                          e.preventDefault();
                          openDialogForRate(mess);
                        }}
                      >
                        Rate
                      </button>
                    </td>
                  )}
                </tr>
              );
            })
          ) : (
            <tr>
              <td
                colSpan={
                  3 +
                  (adminView ? 2 : 0) +
                  (user?.roles?.includes("Student") ? 1 : 0)
                }
                className="py-4 px-4 text-center text-sm text-gray-600"
              >
                Messes are not available...
              </td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};

export default MessTableRate;
