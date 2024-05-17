import { useState, useEffect } from "react";

type SetValue<T> = (value: T | ((val: T) => T)) => void;
type RemoveValue = () => void;
type UseLocalStorageReturn<T> = [T, SetValue<T>, RemoveValue];

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

  const removeValue: RemoveValue = () => {
    localStorage.removeItem(key);
    setStoredValue(initialValue);
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

  return [storedValue, setValue, removeValue];
};

export default useLocalStorage;
