import { useEffect, useState, Dispatch, SetStateAction } from "react";

type SetValue<T> = Dispatch<SetStateAction<T>>;
type UseLocalStorageReturn<T> = [T, SetValue<T>];

const useLocalStorage = <T>(
  key: string,
  initialValue: T
): UseLocalStorageReturn<T> => {
  const [storedValue, setStoredValue] = useState<T>(() => {
    const item = localStorage.getItem(key);
    return item ? (JSON.parse(item) as T) : initialValue;
  });

  const setValue: SetValue<T> = (value) => {
    const valueToStore = value instanceof Function ? value(storedValue) : value;
    setStoredValue(valueToStore);
    localStorage.setItem(key, JSON.stringify(valueToStore));
  };

  useEffect(() => {
    const handleStorageChange = () => {
      const item = localStorage.getItem(key);
      setStoredValue(item ? (JSON.parse(item) as T) : initialValue);
    };

    window.addEventListener("storage", handleStorageChange);

    return () => {
      window.removeEventListener("storage", handleStorageChange);
    };
  }, [initialValue, key]);

  return [storedValue, setValue];
};

export default useLocalStorage;
