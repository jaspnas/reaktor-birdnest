import style from "../drones.module.css"

const Drone = ({droneObject}) => {
    const pilot = droneObject.pilot
    return (
        <li className={style.drone}>
        <div className={style.droneDiv}>
            <h3>Drone details:</h3>
            <p>Serial number: {droneObject.serial_number}</p>
            <p>Last seen: {droneObject.last_seen}</p>
            <p>Closest Distance: {Math.round(droneObject.closest_distance/1000)} m</p>
            <h4 className={style.hdg}>Pilot details:</h4>
            <p>Name: {pilot.name}</p>
            <p>Email: {pilot.email}</p>
            <p>Phone number: {pilot.phone_number}</p>
        </div>
        </li>
    );
};

export default Drone
