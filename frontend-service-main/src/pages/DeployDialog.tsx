// DialogForm.tsx
import React, { useState } from 'react';

interface DialogProps {
  isOpen: boolean;
  repo_scout_id: string;
  name: string;
  image: string;
  onClose: () => void;
  onSubmit: (formData: {
    repo_scout_id: string;
    container_port: number;
    image: string;
    name: string;
    replicas: number;
  }) => void;
}

const DeployDialog: React.FC<DialogProps> = ({ isOpen, onClose, repo_scout_id, name, image, onSubmit }) => {
  const [containerPort, setContainerPort] = useState<number>(0);
  const [replicas, setReplicas] = useState<number>(1);
  const [deploymentName, setDeploymentName] = useState<string>(name);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({
      repo_scout_id,
      container_port: Number(containerPort),
      image,
      name: deploymentName,
      replicas,
    });
    onClose(); // Close the dialog after submission
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-75 z-50">
      <div className="bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
        <h2 className="text-xl font-semibold mb-4">Create a Deployment</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-gray-700 font-medium mb-2" htmlFor="field1">
              Repo Scout ID
            </label>
            <input
              type="text"
              id="field1"
              value={repo_scout_id}
              readOnly
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 font-medium mb-2" htmlFor="field2">
              Image
            </label>
            <input
              type="text"
              id="field2"
              value={image}
              readOnly
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 blur-sm"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 font-medium mb-2" htmlFor="field3">
              Name
            </label>
            <input
              type="text"
              id="field3"
              value={deploymentName}
              onChange={(e) => setDeploymentName(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700">Container Port</label>
            <input
              type="text"
              value={containerPort}
              onChange={(e) => setContainerPort(Number(e.target.value))}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700">Replicas</label>
            <input
              type="number"
              value={replicas}
              onChange={(e) => setReplicas(Number(e.target.value))}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>
          <div className="flex justify-end">
            <button
              type="button"
              onClick={onClose}
              className="mr-2 py-2 px-4 bg-gray-500 text-white rounded-md hover:bg-gray-600"
            >
              Cancel
            </button>
            <button
              type="submit"
              className="py-2 px-4 bg-blue-600 text-white rounded-md hover:bg-blue-700"
            >
              Deploy
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default DeployDialog;
