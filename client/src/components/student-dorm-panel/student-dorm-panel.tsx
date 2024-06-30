import React, { useEffect, useState } from "react";
import QRCode from "react-qr-code";
import { Room } from "../../models/room.model";
import { Dorm } from "../../models/dorm.model";
import { useSelector } from "react-redux";
import { RootState } from "../../redux/store/store";
import {
  castFromApplicationTypeNumberToActualString,
  castFromToaletTypeNumberToActualString,
} from "../../utils/converter.utils";

const StudentDormPanel: React.FC<{
  room: Room;
  dorm: Dorm;
}> = ({ room, dorm }) => {
  const user = useSelector((state: RootState) => state.user.user);
  const [qrCodeFormattedValue, setQrCodeFormattedValue] = useState<string>("");
  useEffect(() => {
    setQrCodeFormattedValue(
      room.id + "," + dorm.id + "," + user?.personalIdentificationNumber
    );
  }, []);

  return (
    <div className="flex flex-row items-center justify-center w-full">
      <div className="bg-gunmetal-500 p-5 flex flex-col items-center justify-center">
        <QRCode
          size={256}
          style={{ height: "auto", maxWidth: "100%", width: "100%" }}
          value={qrCodeFormattedValue}
          viewBox={`0 0 256 256`}
        />
        <p className="text-2xl text-white py-3 text-center">
          Scan this QR code to use it
        </p>
      </div>
      <div className="bg-white h-full w-full">
        <div className="flex flex-row">
          <div className="flex flex-col w-1/2">
            <h3 className="text-3xl text-center">Dorm details</h3>
            <ul className="list-none pl-3">
              <li className="mb-2">
                Dorm name: <span className="font-bold">{dorm.name}</span>
              </li>
              <li className="mb-2">
                Dorm location:{" "}
                <span className="font-bold">{dorm.location}</span>
              </li>
              <li className="mb-2">
                <span className="font-bold">Prices:</span>
              </li>
              <ul className="list-disc pl-4">
                {dorm.prices.map((price) => (
                  <li className="mb-2">
                    {castFromApplicationTypeNumberToActualString(
                      price.applicationType
                    )}
                    : <span className="font-bold">{price.price}</span>
                  </li>
                ))}
              </ul>
            </ul>
          </div>
          <div className="flex flex-col w-1/2">
            <h3 className="text-3xl text-center">Room details</h3>
            <ul className="list-none pl-3">
              <li className="mb-2">
                Square foot of room:{" "}
                <span className="font-bold">{room.squareFoot}</span>
              </li>
              <li className="mb-2">
                Number of beds:{" "}
                <span className="font-bold">{room.numberOfBeds}</span>
              </li>
              <li className="mb-2">
                Number of beds:{" "}
                <span className="font-bold">{room.numberOfBeds}</span>
              </li>
              <li className="mb-2">
                Toalet:{" "}
                <span className="font-bold">
                  {castFromToaletTypeNumberToActualString(room.toalet)}
                </span>
              </li>
              <li className="mb-2">
                <span className="font-bold">Roommates:</span>
              </li>
              <ul className="list-disc pl-4">
                {room.students.map((student) => (
                  <li className="mb-2">{student.fullName}</li>
                ))}
              </ul>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
};

export default StudentDormPanel;
