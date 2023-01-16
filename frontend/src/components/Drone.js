import style from "../Drones.module.css"

const Drone = ({droneObject}) => {
    const pilot = droneObject.pilot
    return (
        <li className={style.drone}>
        <div className={style.drone}>
            <h3>Drone details:</h3>
            <p>Serial number: {droneObject.serial_number}</p>
            <p>Last seen: {droneObject.last_seen}</p>
            <p>Closest Distance: {Math.round(droneObject.closest_distance/1000)} m</p>
            <h5>Pilot details:</h5>
            <p>Name: {pilot.name}</p>
            <p>Email: {pilot.email}</p>
            <p>Phone number: {pilot.phone_number}</p>
        </div>
        </li>
    );
};

export default Drone
