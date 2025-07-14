import { useState } from "react";

export default function ImageUploader() {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files) {
      setSelectedFile(event.target.files[0]);
    }
  };

  const handleUpload = async () => {
    if (!selectedFile) return;

    const formData = new FormData();

    formData.append("image", selectedFile);

    try {
      const response = await fetch("api/analyze", {
        method: "POST",
        body: formData,
      });
      const result = await response.json();
      console.log(result.message);
    } catch (error) {
      console.error("Error: ", error);
    }
  };

  return (
    <>
      <h2>Upload Image:</h2>
      <input type="file" accept="image/*" onChange={handleFileChange} />

      <button onClick={handleUpload} disabled={!selectedFile}>
        Upload
      </button>
    </>
  );
}
