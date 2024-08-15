import React, { useEffect, useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import { University } from "../../models/university.model";
import { axiosInstance } from "../../services/axios.service";

const AllUniversities = () => {
  const [universities, setUniversities] = useState<University[]>([]);


  useEffect(() => {
    axiosInstance
      .get("/university/")
      .then((response) => {
        console.log(response); 
        setUniversities(response.data.data);
      })
      .catch((err) => {
        console.error("Error fetching universities:", err);
      });
  }, []);


  return (
    <div className="overflow-x-auto">
      <table className="min-w-full bg-white">
        <thead>
          <tr>
            <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
              Name
            </th>
            <th className="py-2 px-4 border-b border-gray-300 text-left text-sm font-medium text-gray-700">
              Address
            </th>
          </tr>
        </thead>
        <tbody>
          {universities && universities.length > 0 ? (
            universities.map((uni) => (
              <tr key={uni.id || uni.name}>
                <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                  {uni.name}
                </td>
                <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                  {uni.address}
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td className="py-2 px-4 border-b border-gray-300 text-sm text-center">
                No universities found.
              </td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};

export default AllUniversities;
