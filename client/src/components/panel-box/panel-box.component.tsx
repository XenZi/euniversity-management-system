import React, { FC } from "react";

interface PanelBoxProps {
  onClick?: () => void;
  panelBoxTitle?: string;
  panelBoxDescription?: string;
}
const PanelBox: FC<PanelBoxProps> = ({
  onClick,
  panelBoxTitle,
  panelBoxDescription,
}) => {
  return (
    <div
      className="bg-battleship-400 mx-2 text-white p-5 cursor-pointer"
      onClick={onClick}
    >
      {panelBoxDescription}
      {panelBoxTitle}
    </div>
  );
};

export default PanelBox;
