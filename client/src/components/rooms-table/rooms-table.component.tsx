import React, { ChangeEvent, useEffect, useState } from "react";
import { Room } from "../../models/room.model";
import { Dorm } from "../../models/dorm.model";
import { axiosInstance } from "../../services/axios.service";
import { useDispatch } from "react-redux";
import { useModalContext } from "../../context/modal.context";
import { closeModal, setModalOpen } from "../../redux/slices/modal.slice";
import DeleteDialog from "../dialogs/delete-dialog/delete-dialog.component";
import EditRoomForm from "../forms/edit-room/edit-room.form";

const RoomsTable = () => {
  const [loadedRooms, setLoadedRooms] = useState<Room[]>([]);
  const [dormList, setDormList] = useState<Dorm[]>([]);
  const [selectedDorm, setSelectedDorm] = useState<Dorm>();
  const dispatch = useDispatch();
  const { setContent } = useModalContext();

  const handleSelectChange = (e: ChangeEvent<HTMLSelectElement>) => {
    const selectedId = e.target.value;
    const dorm = dormList.find((d) => d.id === selectedId);
    setSelectedDorm(dorm);
  };

  useEffect(() => {
    axiosInstance
      .get(`/dorm/${selectedDorm?.id}/rooms`)
      .then((data) => {
        setLoadedRooms(data.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
    console.log(selectedDorm);
  }, [selectedDorm]);
  useEffect(() => {
    axiosInstance
      .get("/dorm/all")
      .then((data) => {
        setDormList(data.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const openDialogForDelete = (roomID: string) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    setContent(
      <DeleteDialog
        functionToProceedOnDelete={() => {
          deleteRoom(roomID);
        }}
      />
    );
  };

  const openDialogForEdit = (room: Room) => {
    dispatch(closeModal());
    dispatch(setModalOpen());
    setContent(<EditRoomForm room={room} />);
  };

  const deleteRoom = (roomID: string) => {
    axiosInstance
      .delete(`/dorm/room/${roomID}`)
      .then((data) => {
        console.log(data);
      })
      .catch((err) => {
        console.log(err);
      });
    dispatch(closeModal());
  };
  return (
    <>
      <form className="flex flex-col">
        <label htmlFor="options">Choose an dorm:</label>
        <select
          id="options"
          value={selectedDorm?.id}
          onChange={handleSelectChange}
        >
          <option value="" disabled selected>
            Select an option
          </option>
          {dormList.map((dorm, index) => (
            <option key={index} value={dorm.id}>
              {dorm.name} - {dorm.location}
            </option>
          ))}
        </select>
      </form>
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white">
          <thead>
            <tr>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Square Foot
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Toalet Type
              </th>
              <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
                Number of Beds in Room
              </th>
            </tr>
          </thead>
          <tbody>
            {loadedRooms && loadedRooms.length > 0 ? (
              loadedRooms.map((room, index) => (
                <tr key={index}>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                    {room.squareFoot}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                    {room.toalet}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                    {room.numberOfBeds}
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                    <button
                      className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                      onClick={(e) => {
                        e.preventDefault();
                        openDialogForEdit(room);
                      }}
                    >
                      Edit room
                    </button>
                  </td>
                  <td className="py-2 px-4 border-b border-gray-300 text-sm ">
                    <button
                      className="border bg-auburn-500 border-auburn-500 font-semibold py-1 px-2 rounded focus:border-auburn-700 text-white"
                      onClick={(e) => {
                        e.preventDefault();
                        openDialogForDelete(room.id);
                      }}
                    >
                      Delete room
                    </button>
                  </td>
                </tr>
              ))
            ) : (
              <tr>
                <td
                  colSpan={3}
                  className="py-2 px-4 border-b border-gray-300 text-sm text-center"
                >
                  No rooms available
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </>
  );
};

export default RoomsTable;
