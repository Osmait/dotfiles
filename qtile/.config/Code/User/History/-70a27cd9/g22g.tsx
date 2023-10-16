import {
  faCamera,
  faCameraAlt,
  faUser,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import axios from "axios";
import React from "react";
import { useDropzone } from "react-dropzone";

const API = process.env.NEXT_PUBLIC_API;

export function DropZone({ id }: any) {
  const { acceptedFiles, getRootProps, getInputProps } = useDropzone({
    maxFiles: 1,
  });

  const files = acceptedFiles.map((file: any) => (
    <li key={file.path}>
      {file.path} - {file.size} bytes
    </li>
  ));

  const handlerSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const formData = new FormData();
    formData.append("file", acceptedFiles[0]);

    console.log(formData);
    try {
      axios.post(`${API}/client/upload/${id}`, formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <form onSubmit={handlerSubmit}>
      <section className="container">
        <div {...getRootProps({ className: "dropzone" })}>
          <input {...getInputProps()} />
          <p>Subir Imagen</p>
          <p>{files}</p>
        </div>
      </section>
      <button type="submit">Subir</button>
    </form>
  );
}
