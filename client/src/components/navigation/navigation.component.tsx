import React from "react";
import { Link } from "react-router-dom";

const Navigation = () => {
  const navLinks = [
    {
      url: "/home",
      name: "Home",
    },
    {
      url: "/dorm",
      name: "Dorm",
    },
  ];
  return (
    <div className="bg-papaya-500 w-full p-3">
      <div className="max-w-7xl mx-auto px-8">
        <div className="flex items-center justify-between h-16">
          <Link to={"/"}>EUniversity</Link>
          <div className="flex space-x-4">
            {navLinks.map((link, i) => {
              return (
                <Link
                  to={link.url}
                  key={i}
                  className="px-3 py-3 hover:text-teal-500"
                >
                  {link.name}
                </Link>
              );
            })}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Navigation;
