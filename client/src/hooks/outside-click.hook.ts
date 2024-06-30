import { useEffect, RefObject } from "react";

const useOutsideClick = (
  ref: RefObject<HTMLElement>,
  callback: () => void,
  isActive: boolean
) => {
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (ref.current && !ref.current.contains(event.target as Node)) {
        callback();
      }
    };

    if (isActive) {
      document.addEventListener("mousedown", handleClickOutside);
    } else {
      document.removeEventListener("mousedown", handleClickOutside);
    }

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [ref, callback, isActive]);
};

export default useOutsideClick;
