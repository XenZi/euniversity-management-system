// ModalContext.tsx
import React, { createContext, useState, useContext, ReactNode } from "react";

interface ModalContextType {
  content: ReactNode;
  setContent: (content: ReactNode) => void;
}

const ModalContext = createContext<ModalContextType | undefined>(undefined);

export const ModalProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [content, setContent] = useState<ReactNode>(null);

  return (
    <ModalContext.Provider value={{ content, setContent }}>
      {children}
    </ModalContext.Provider>
  );
};

export const useModalContext = (): ModalContextType => {
  const context = useContext(ModalContext);
  if (!context) {
    throw new Error("useModalContext must be used within a ModalProvider");
  }
  return context;
};
