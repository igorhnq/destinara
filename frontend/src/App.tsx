import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Welcome from './pages/Welcome'
import Login from './pages/Login'
import Register from './pages/Register'
import TravelPackages from './pages/TravelPackages'
import PackageDetails from './pages/PackageDetails'

export default function App() {
	return (
		<Router>
			<Routes>
				<Route path="/" element={<Welcome />} />
				<Route path="/login" element={<Login />} />
				<Route path="/register" element={<Register />} />
				<Route path="/travel-packages" element={<TravelPackages />} />
				<Route path="/package-details" element={<PackageDetails />} />
			</Routes>
		</Router>
	)
}