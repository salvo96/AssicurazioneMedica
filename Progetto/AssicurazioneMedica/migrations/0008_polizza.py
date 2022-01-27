# Generated by Django 4.0 on 2021-12-17 14:23

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('AssicurazioneMedica', '0007_alter_datimedici_fumatore_and_more'),
    ]

    operations = [
        migrations.CreateModel(
            name='Polizza',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('data_sottoscrizione', models.DateField()),
                ('scadenza', models.DateField()),
                ('premio_assicurativo', models.FloatField()),
                ('cliente', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='AssicurazioneMedica.cliente')),
            ],
        ),
    ]