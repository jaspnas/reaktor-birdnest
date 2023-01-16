import React, {useEffect, useRef, useState} from "react";
import axios from "axios";
import Drone from "./Drone.js"
import styles from "../drones.module.css"


const updateData = (data, update) => {
    let updated = false
    let newdata = data.map(element => {
        if (update.serial_number === element.serial_number) {
            if (update.closest_distance > element.closest_distance) {
                update.closest_distance = element.closest_distance
            }
            updated = true
            return update
        } else return element
    })
    newdata = newdata.filter((element) => { return !((new Date() - new Date(element.last_seen)) > (10 * 60 * 1000))})
    if (!updated) newdata = newdata.concat(update)
    return newdata
}


const socket = new WebSocket("wss://" + process.env.REACT_APP_WEBSITE_DOMAIN + "/api/ws")

const DroneList = () => {
    const domain = process.env.REACT_APP_WEBSITE_DOMAIN
    let [data, setData] = useState([])
    useEffect(() => {
        axios.get("https://" + domain + "/api/drones").then(resp => {
            const data = resp.data
            setData(data)
        })
    }, [])

    const ws = useRef(null)

    useEffect(() => {

        socket.onopen = () => {
            console.log("Connected to websocket")
        }

        socket.onmessage = (e) => {
            setData(updateData(data, JSON.parse(e.data)))
        }

        socket.onerror = (err) => {
            console.log("Error: ", err)
        }

        ws.current = socket

    }, [data])

    if (!data) return (<div>Loading...</div>)
    return (
        <div>
            <ul className={styles.droneList}>
                {data.map((drone) => {
                    return (
                        <Drone key={drone.serial_number} droneObject={drone}/>
                    )
                })
                }
            </ul>
        </div>
    )
}

export default DroneList