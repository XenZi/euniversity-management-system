import React, { useEffect, useState } from "react";
import QRCode from "react-qr-code";
import { Room } from "../../models/room.model";
import { Dorm } from "../../models/dorm.model";
import { useSelector } from "react-redux";
import { RootState } from "../../redux/store/store";

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
    <div className="flex flex-row items-center justify-center h-5/6">
      <div className="bg-gunmetal-500 p-5 flex flex-col items-center justify-center">
        <QRCode
          size={256}
          style={{ height: "auto", maxWidth: "100%", width: "100%" }}
          value={qrCodeFormattedValue}
          viewBox={`0 0 256 256`}
        />
        <p className="text-2xl text-white py-3">Scan this QR code to use it</p>
      </div>
      <div className="bg-white"></div>
    </div>
  );
};

export default StudentDormPanel;
