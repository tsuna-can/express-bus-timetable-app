package com.tsunacan.expressbustimetableapp.data

import androidx.datastore.core.CorruptionException
import androidx.datastore.core.Serializer
import com.google.protobuf.InvalidProtocolBufferException
import com.tsunacan.expressbustimetableapp.DefaultBusStop
import java.io.InputStream
import java.io.OutputStream
import javax.inject.Inject

class UserSettingsSerializer @Inject constructor(): Serializer<DefaultBusStop> {
    override val defaultValue: DefaultBusStop = DefaultBusStop.getDefaultInstance()
    override suspend fun readFrom(input: InputStream): DefaultBusStop {
        try {
            return DefaultBusStop.parseFrom(input)
        } catch (exception: InvalidProtocolBufferException) {
            throw CorruptionException("Cannot read proto.", exception)
        }
    }

    override suspend fun writeTo(t: DefaultBusStop, output: OutputStream) = t.writeTo(output)
}